package policy

import (
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	policy "github.com/CodeLine-95/go-cloud-native/internal/app/service/casbin"
)

var e *policy.CasbinService

func Init() {
	// 初始化 policy 做持久化
	e = &policy.CasbinService{
		Type:      policy.RBAC_DEFAULT,
		DB:        db.D(),
		Prefix:    "cloud",
		TableName: "casbin",
	}
	e.Init()
}

// 获取初始化后的句柄
func Casbin() *policy.CasbinService {
	return e
}
