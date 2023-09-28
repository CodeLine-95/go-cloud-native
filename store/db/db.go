package db

import (
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/pkg/logz"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"sync"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var (
	groups map[string]*xorm.Engine
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
	}
)

func Init() {
	once.Do(func() {
		groups = make(map[string]*xorm.Engine)
		var cfgs map[string]cfg
		if err := viper.UnmarshalKey("db", &cfgs); err != nil {
			panic(err)
		}
		logz.Warn(fmt.Sprintln(cfgs))
		for instanceRwType, instanceCfg := range cfgs {
			dsn := fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=%s",
				instanceCfg.User,
				instanceCfg.Pass,
				instanceCfg.Host,
				instanceCfg.Port,
				instanceCfg.Name,
				"utf8mb4")
			engine, err := xorm.NewEngine("mysql", dsn)
			if err != nil {
				panic(err)
			}
			if instanceCfg.MaxIdleConn == 0 {
				instanceCfg.MaxIdleConn = 5
			}
			engine.SetMaxIdleConns(instanceCfg.MaxIdleConn)
			if instanceCfg.MaxOpenConn == 0 {
				instanceCfg.MaxOpenConn = 10
			}
			engine.SetMaxOpenConns(instanceCfg.MaxOpenConn)
			engine.ShowSQL(viper.GetBool("app.debug"))
			engine.SetMapper(names.GonicMapper{})

			groups[instanceRwType] = engine
		}
	})
}

// Grp 返回指定实例组实例
func Grp(name string) *xorm.Engine {
	return groups[name]
}
