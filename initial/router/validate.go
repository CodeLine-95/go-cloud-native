package router

import (
	"github.com/CodeLine-95/go-cloud-native/internal/app/logic"
	"github.com/gin-gonic/gin"
)

// 入口
func Validate(r *gin.RouterGroup) {
	userRouter(r)
	roleRouter(r)
	menuRouter(r)
	allocationRouter(r)
	dockerRouter(r)
	etcdRouter(r)
}

func userRouter(r *gin.RouterGroup) {
	c := r.Group("/user")
	c.GET("info", logic.GetUserInfo)
}

// 角色
func roleRouter(r *gin.RouterGroup) {
	c := r.Group("/role")
	c.POST("/list", logic.RoleResp)
	c.POST("/add", logic.RoleAdd)
	c.POST("/edit", logic.RoleEdit)
	c.POST("/del", logic.RoleDel)
	c.POST("/menu", logic.GetRoleMenu)
}

// 菜单
func menuRouter(r *gin.RouterGroup) {
	c := r.Group("/menu")
	c.POST("/list", logic.MenuResp)
	c.POST("/add", logic.MenuAdd)
	c.POST("/edit", logic.MenuEdit)
	c.POST("/del", logic.MenuDel)
}

// 分配
func allocationRouter(r *gin.RouterGroup) {
	c := r.Group("/allocation")
	c.POST("/user-role", logic.UserRole)
	c.POST("/role-menu", logic.RoleMenu)
}

// docker
func dockerRouter(r *gin.RouterGroup) {
	docker := r.Group("/docker")
	docker.GET("/list", logic.ContainerList)
	docker.POST("/logs", logic.ContainerLogs)
	docker.POST("/stop", logic.ContainerStop)
	docker.POST("/batch-stop", logic.BatchContainerStop)
	docker.POST("/create", logic.ContainerCreate)
}

// etcd
func etcdRouter(r *gin.RouterGroup) {
	etcd := r.Group("/etcd")
	etcd.POST("/create", logic.CreatePut)
	etcd.GET("/list", logic.GetService)
	etcd.DELETE("/del", logic.DelService)
	etcd.POST("/put", logic.PutService)
}
