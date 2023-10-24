package policy

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"sync"
)

const (
	DIR_PATH      = "tools/casbin"
	CONF_PATH     = "conf"
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

var confMap map[string]string

// 持久化到数据库

var (
	syncecEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// init 自动初始化
func (c *CasbinService) Init() {
	pwd, _ := os.Getwd()
	// 切换到指定目录下
	_ = os.Chdir(pwd + "/" + DIR_PATH)
	// 重新获取当前目录
	pwd, _ = os.Getwd()
	path := pwd + "/" + CONF_PATH
	//获取文件或目录相关信息
	fileInfoList, err := os.ReadDir(filepath.Clean(filepath.ToSlash(path)))
	if err != nil {
		panic(err)
	}
	confMap = make(map[string]string, len(fileInfoList))
	for i := range fileInfoList {
		confMap[fileInfoList[i].Name()] = fileInfoList[i].Name()
	}

	once.Do(func() {
		switch c.Type {
		case RBAC_DEFAULT:
			c.Type = RBAC_DEFAULT
		case RBAC_DOMAINS:
			c.Type = RBAC_DOMAINS
		case RBAC_RESOURCE:
			c.Type = RBAC_RESOURCE
		default:
			c.Type = RBAC_DEFAULT
		}
		// 通过现有的gorm实例和指定的表前缀和表名创建gorm适配器
		a, _ := gormadapter.NewAdapterByDBUseTableName(c.DB, c.Prefix, c.TableName)
		// policy 初始化持久化到 DB
		syncecEnforcer, _ = casbin.NewSyncedEnforcer(path+"/"+confMap[c.Type+".conf"], a)
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
