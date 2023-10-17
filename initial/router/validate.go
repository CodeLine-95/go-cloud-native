package router

import (
	"github.com/CodeLine-95/go-cloud-native/internal/app/logic"
	"github.com/gin-gonic/gin"
)

func Validate(r *gin.RouterGroup) {
	RoleRouter(r)
	MenuRouter(r)
}

func RoleRouter(r *gin.RouterGroup) {
	c := r.Group("/role")
	c.GET("/list", logic.RoleResp)
	c.POST("/add", logic.RoleAdd)
	c.POST("/edit", logic.RoleEdit)
	c.POST("/del", logic.RoleDel)
}

func MenuRouter(r *gin.RouterGroup) {
	c := r.Group("/menu")
	c.GET("/list", logic.MenuResp)
	c.POST("/add", logic.MenuAdd)
	c.POST("/edit", logic.MenuEdit)
	c.POST("/del", logic.MenuDel)
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
