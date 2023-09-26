package routers

import (
	"github.com/CodeLine-95/go-cloud-native/controllers/container"
	"github.com/gin-gonic/gin"
)

func DockerRouter(r *gin.RouterGroup) {
	dockerApi := container.DockerApi{}
	docker := r.Group("/docker")
	docker.GET("/container-list", dockerApi.GetContainerList)
	docker.POST("/container-logs", dockerApi.ContainerLogs)

	docker.GET("/images-list", dockerApi.GetImageList)
	docker.POST("/images-pull", dockerApi.ImagePull)
}
