package container

import (
	"github.com/CodeLine-95/go-cloud-native/pkg/containers"
	"github.com/gin-gonic/gin"
)

// DockerApiInterface docker 容器
type DockerApiInterface interface {
	GetContainerList(c *gin.Context) // 查看容器列表
	ContainerCreate(c *gin.Context)  // 运行容器
	ContainerLogs(c *gin.Context)    // 获取指定容器的日志
	ContainerStop(c *gin.Context)    // 停止运行中的容器 / 停止所有运行中的容器
	GetImageList(c *gin.Context)     // 查看镜像列表
	ImagePull(c *gin.Context)        // 拉取镜像
	ContainerCommit(c *gin.Context)  // 保存容器成镜像
}

type DockerApi struct {
}

func (d DockerApi) GetContainerList(c *gin.Context) {
	docker := containers.Docker{}

	containerList, err := docker.ContainerList()
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": containerList,
	})

	return
}
