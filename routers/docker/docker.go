package docker

import (
	"github.com/CodeLine-95/go-cloud-native/controllers/container"
	"github.com/gin-gonic/gin"
)

func RouterDocker(r *gin.RouterGroup) {
	dockerApi := container.DockerApi{}
	docker := r.Group("/docker")
	docker.GET("/container-list", dockerApi.GetContainerList)
}
