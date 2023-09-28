package container

import (
	"github.com/CodeLine-95/go-cloud-native/services"
	"github.com/CodeLine-95/go-cloud-native/services/containers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

func (d DockerApi) ContainerStop(c *gin.Context) {
	docker := containers.Docker{}
	var param services.ContainerStop
	if err := c.ShouldBindJSON(&param); err != nil {
		panic(err)
	}
	var containerID []string

	containerID = strings.Split(param.Ids, ",")
	code, codeErr := docker.ContainerStop(containerID)
	if codeErr != nil {
		panic(codeErr)
	}

	c.JSON(code, gin.H{})

	return
}

func (d DockerApi) ContainerLogs(c *gin.Context) {
	docker := containers.Docker{}
	var param services.ContainerStop
	if err := c.ShouldBindJSON(&param); err != nil {
		panic(err)
	}

	out := docker.ContainerLogs(param.Ids)
	c.JSON(http.StatusOK, gin.H{
		"data": out,
	})

	return
}

func (d DockerApi) GetImageList(c *gin.Context) {
	docker := containers.Docker{}

	imageList, err := docker.ImageList()
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": imageList,
	})

	return
}

func (d DockerApi) ImagePull(c *gin.Context) {
	docker := containers.Docker{}
	var param services.ImagesPull
	if err := c.ShouldBindJSON(&param); err != nil {
		panic(err)
	}

	out := docker.ImagePull(param.Refstr)
	c.JSON(http.StatusOK, gin.H{
		"data": out,
	})

	return
}
