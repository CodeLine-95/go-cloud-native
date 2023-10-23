package policy

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"os"
	"sync"
)

const (
	RBAC_DEFAULT  = "rbac_default"
	RBAC_DOMAINS  = "rbac_domains"
	RBAC_RESOURCE = "rbac_resource_roles"
)

type CasbinService struct {
	Type      string   // 规则类型
	DB        *gorm.DB // 数据库句柄
	Prefix    string   // 自定义表前缀
	TableName string   // 自定义表名
}

var CasbinServiceApp = new(CasbinService)

var confMap map[string]string

// 持久化到数据库

var (
	syncecEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// init 自动初始化
func init() {
	pwd, _ := os.Getwd()
	//获取文件或目录相关信息
	fileInfoList, err := os.ReadDir(pwd + "/conf")
	if err != nil {
		panic(err)
	}
	for i := range fileInfoList {
		confMap[fileInfoList[i].Name()] = fileInfoList[i].Name()
	}

	once.Do(func() {
		// 通过现有的gorm实例和指定的表前缀和表名创建gorm适配器
		a, _ := gormadapter.NewAdapterByDBUseTableName(CasbinServiceApp.DB, CasbinServiceApp.Prefix, CasbinServiceApp.TableName)
		// policy 初始化持久化到 DB
		syncecEnforcer, _ = casbin.NewSyncedEnforcer(confMap[CasbinServiceApp.Type], a)
	})

	// 从DB中 load 策略
	err = syncecEnforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}
}

// Checked 策略验证 sub(角色), obj(路径), act(方法)
func (c *CasbinService) Checked(rivals ...interface{}) bool {
	ok, err := syncecEnforcer.Enforce(rivals)
	if err != nil {
		return false
	}
	return ok
}

// AddPolicy 添加策略 sub(角色), obj(路径), act(方法)
func (c *CasbinService) AddPolicy(rivals ...interface{}) bool {
	ok, err := syncecEnforcer.AddPolicy(rivals)
	if err != nil {
		return false
	}
	return ok
}

// RemovePolicy 删除策略 sub(角色), obj(路径), act(方法)
func (c *CasbinService) RemovePolicy(rivals ...interface{}) bool {
	ok, err := syncecEnforcer.RemovePolicy(rivals)
	if err != nil {
		return false
	}
	return ok
}
