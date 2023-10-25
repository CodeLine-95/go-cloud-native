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
}

// 菜单
func menuRouter(r *gin.RouterGroup) {
	c := r.Group("/menu")
	c.POST("/list", logic.MenuResp)
	c.POST("/add", logic.MenuAdd)
	c.POST("/edit", logic.MenuEdit)
	c.POST("/del", logic.MenuDel)
	c.POST("/role-list", logic.GetRoleMenu)
}

// 分配
func allocationRouter(r *gin.RouterGroup) {
	c := r.Group("/allocation")
	c.POST("/user-role", logic.UserRole)
	c.POST("/role-menu", logic.RoleMenu)
}

//func DockerRouter(r *gin.RouterGroup) {
//	dockerApi := logic.DockerApi{}
//	docker := r.Group("/docker")
//	docker.GET("/container-list", dockerApi.GetContainerList)
//	docker.POST("/container-logs", dockerApi.ContainerLogs)
//
//	docker.GET("/images-list", dockerApi.GetImageList)
//	docker.POST("/images-pull", dockerApi.ImagePull)
//}
