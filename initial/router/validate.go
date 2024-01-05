package router

import (
	"github.com/CodeLine-95/go-cloud-native/internal/app/logic"
	"github.com/gin-gonic/gin"
)

// 入口
func Validate(r *gin.RouterGroup) {
	systemRouter(r)
	userRouter(r)
	roleRouter(r)
	menuRouter(r)
	assignRouter(r)
	dockerRouter(r)
	etcdRouter(r)
}

func userRouter(r *gin.RouterGroup) {
	c := r.Group("/user")
	{
		c.GET("list", logic.GetUserList)
		c.GET("info", logic.GetUserInfo)
	}
}

func systemRouter(r *gin.RouterGroup) {
	c := r.Group("/log")
	{
		c.GET("list", logic.GetLogList)
	}
}

// 角色
func roleRouter(r *gin.RouterGroup) {
	c := r.Group("/role")
	{
		c.GET("/list", logic.RoleResp)
		c.POST("/add", logic.RoleAdd)
		c.PUT("/edit", logic.RoleEdit)
		c.DELETE("/del", logic.RoleDel)
		c.GET("/menu", logic.GetRoleMenu)
	}
}

// 菜单
func menuRouter(r *gin.RouterGroup) {
	c := r.Group("/menu")
	{
		c.GET("/list", logic.MenuResp)
		c.POST("/add", logic.MenuAdd)
		c.PUT("/edit", logic.MenuEdit)
		c.DELETE("/del", logic.MenuDel)
	}
}

// 分配
func assignRouter(r *gin.RouterGroup) {
	c := r.Group("/assign")
	{
		c.POST("/user-role", logic.UserRole)
		c.POST("/role-menu", logic.RoleMenu)
	}
}

// docker
func dockerRouter(r *gin.RouterGroup) {
	c := r.Group("/docker")
	{
		c.GET("/list", logic.ContainerList)
		c.POST("/logs", logic.ContainerLogs)
		c.POST("/stop", logic.ContainerStop)
		c.POST("/batch-stop", logic.BatchContainerStop)
		c.POST("/create", logic.ContainerCreate)
	}
}

// etcd
func etcdRouter(r *gin.RouterGroup) {
	c := r.Group("/etcd")
	{
		c.POST("/create", logic.CreatePut)
		c.GET("/list", logic.GetService)
		c.DELETE("/del", logic.DelService)
		c.POST("/put", logic.PutService)
	}
}
