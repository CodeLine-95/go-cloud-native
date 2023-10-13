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
	c.GET("/list", logic.List)
	c.POST("/add", logic.Add)
	c.POST("/edit", logic.Edit)
	c.POST("/del", logic.Del)
}

func MenuRouter(r *gin.RouterGroup) {
	//c := r.Group("/menu")
	//c.GET("/list")
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
