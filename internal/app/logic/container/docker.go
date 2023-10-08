package container

import (
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/app/service/containers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type DockerApi struct {
}

// GetContainerList 查看容器列表
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

// ContainerStop 停止运行中的容器 / 停止所有运行中的容器
func (d DockerApi) ContainerStop(c *gin.Context) {
	docker := containers.Docker{}
	var param models.ContainerStopRequest
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

// ContainerLogs 获取指定容器的日志
func (d DockerApi) ContainerLogs(c *gin.Context) {
	docker := containers.Docker{}
	var param models.ContainerStopRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		panic(err)
	}

	out := docker.ContainerLogs(param.Ids)
	c.JSON(http.StatusOK, gin.H{
		"data": out,
	})

	return
}

// GetImageList 查看镜像列表
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

// ImagePull 拉取镜像
func (d DockerApi) ImagePull(c *gin.Context) {
	docker := containers.Docker{}
	var param models.ImagesPullRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		panic(err)
	}

	out := docker.ImagePull(param.Refstr)
	c.JSON(http.StatusOK, gin.H{
		"data": out,
	})

	return
}
