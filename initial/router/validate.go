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
	assignRouter(r)
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
	c.GET("/list", logic.RoleResp)
	c.POST("/add", logic.RoleAdd)
	c.PUT("/edit", logic.RoleEdit)
	c.DELETE("/del", logic.RoleDel)
	c.GET("/menu", logic.GetRoleMenu)
}

// 菜单
func menuRouter(r *gin.RouterGroup) {
	c := r.Group("/menu")
	c.GET("/list", logic.MenuResp)
	c.POST("/add", logic.MenuAdd)
	c.PUT("/edit", logic.MenuEdit)
	c.DELETE("/del", logic.MenuDel)
}

// 分配
func assignRouter(r *gin.RouterGroup) {
	c := r.Group("/assign")
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
