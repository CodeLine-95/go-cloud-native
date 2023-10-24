package db

import (
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

var (
	groups map[string]*gorm.DB
	once   sync.Once
)

type (

	// group map[rwType]*xorm.Engine，key为"r"或"w"
	// group["r"]为只读实例，group["w"]为写实例
	cfg struct {
		Host        string
		Port        int
		User        string
		Pass        string
		Name        string
		MaxIdleConn int
		MaxOpenConn int
		AutoLoad    bool
		Debug       bool
	}
)

func Init() {
	once.Do(func() {
		groups = make(map[string]*gorm.DB)
		var cfgs map[string]cfg
		if err := viper.UnmarshalKey("db", &cfgs); err != nil {
			logz.Error("get db config reader fail", logz.F("err", err))
			panic(err)
		}
		logz.Info(fmt.Sprintln(cfgs))
		for instanceRwType, instanceCfg := range cfgs {
			dsn := fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=%s",
				instanceCfg.User,
				instanceCfg.Pass,
				instanceCfg.Host,
				instanceCfg.Port,
				instanceCfg.Name,
				"utf8mb4")

			engineConfig := &gorm.Config{
				// 跳过默认事务
				SkipDefaultTransaction: true,
				// 命名策略
				NamingStrategy: nil,
				// 生成 SQL 但不执行，可以用于准备或测试生成的 SQL
				DryRun: false,
				// PreparedStmt 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
				PrepareStmt: true,
				// 在完成初始化后，GORM 会自动 ping 数据库以检查数据库的可用性，若要禁用该特性，可将其设置为 true
				DisableAutomaticPing: false,
				// 在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为
				DisableForeignKeyConstraintWhenMigrating: false,
				// 启用全局更新
				AllowGlobalUpdate: false,
				// 翻译方言错误
				// 如果您希望将数据库的方言错误转换为gorm的错误类型（例如将MySQL中的“Duplicate entry”转换为ErrDuplicatedKey），则在打开数据库连接时启用TranslateError标志。
				TranslateError: true,
			}
			// 打印全部SQL
			if instanceCfg.Debug {
				engineConfig.Logger = logger.Default.LogMode(logger.Info)
			}
			// 链接数据库
			engine, err := gorm.Open(mysql.Open(dsn), engineConfig)
			if err != nil {
				logz.Error("new engine fail", logz.F("instanceRwType", instanceRwType), logz.F("err", err))
				panic(err)
			}

			// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
			sqlDB, err := engine.DB()
			if err != nil {
				logz.Error("sql DB fail", logz.F("err", err))
				panic(err)
			}

			// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
			if instanceCfg.MaxIdleConn == 0 {
				instanceCfg.MaxIdleConn = 5
			}
			sqlDB.SetMaxIdleConns(instanceCfg.MaxIdleConn)

			// SetMaxOpenConns 设置打开数据库连接的最大数量。
			if instanceCfg.MaxOpenConn == 0 {
				instanceCfg.MaxOpenConn = 10
			}
			sqlDB.SetMaxOpenConns(instanceCfg.MaxOpenConn)

			// SetConnMaxLifetime 设置了连接可复用的最大时间。
			sqlDB.SetConnMaxLifetime(time.Hour)

			// 是否自动同步数据库表结构
			if instanceCfg.AutoLoad {
				syncMap := make([]interface{}, 0)
				syncMap = append(
					syncMap,
					models.CloudUser{},
					models.CloudRole{},
					models.CloudRoleMenu{},
					models.CloudMenu{},
					models.CloudUserRole{},
				)
				err := engine.AutoMigrate(syncMap...)
				if err != nil {
					logz.Error("engine sync fail",
						logz.F("table", []string{
							new(models.CloudUser).TableName(),
							new(models.CloudRole).TableName(),
							new(models.CloudRoleMenu).TableName(),
							new(models.CloudMenu).TableName(),
							new(models.CloudUserRole).TableName(),
						}), logz.F("err", err),
					)
				}
			}

			groups[instanceRwType] = engine
		}
	})
}

// Grp 返回指定实例组实例
func Grp(name string) *gorm.DB {
	return groups[name]
}

// D 返回指定实例组实例
func D() *gorm.DB {
	return groups[constant.CloudNative]
}
