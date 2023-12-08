package logic

import (
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/service/containers"
	"github.com/CodeLine-95/go-cloud-native/internal/app/service/containers/docker"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
)

var dockerClient *docker.DockerClient

func init() {
	dockerClient = containers.NewDockerClient()
}

func ContainerList(c *gin.Context) {
	list, err := dockerClient.ContainerList()
	if err != nil {
		response.Error(c, constant.ErrorContainerList, err, constant.ErrorMsg[constant.ErrorContainerList])
		return
	}
	response.OK(c, list, constant.ErrorMsg[constant.Success])
}

func ContainerLogs(c *gin.Context) {
	var params common.ContainerLogsRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	out, err := dockerClient.ContainerLogs(params.ID)
	if err != nil {
		response.Error(c, constant.ErrorContainerLogs, err, constant.ErrorMsg[constant.ErrorContainerLogs])
		return
	}
	response.OK(c, out, constant.ErrorMsg[constant.Success])
}

func ContainerStop(c *gin.Context) {
	var params common.ContainerLogsRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	err := dockerClient.ContainerStop(params.ID)
	if err != nil {
		response.Error(c, constant.ErrorContainerStop, err, constant.ErrorMsg[constant.ErrorContainerStop])
		return
	}
	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

func BatchContainerStop(c *gin.Context) {
	var params common.BatchContainerLogsRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	success, fail := dockerClient.BatchContainerStop(params.ID)
	data := gin.H{
		"success_count": success,
		"fail_count":    fail,
	}
	response.OK(c, data, constant.ErrorMsg[constant.Success])
}

func ContainerCreate(c *gin.Context) {
	var params common.ContainerCreateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	// 配置要启动的容器
	containerOptions := &container.Config{
		Image: params.Image,
		Cmd:   params.Cmd,
	}

	// 主机配置
	hostOptions := &container.HostConfig{
		// 将容器的80端口映射到宿主机的8080端口
		PortBindings: nat.PortMap{
			nat.Port(params.LocalProt + "/tcp"): []nat.PortBinding{
				{
					HostIP:   params.HostIP,
					HostPort: params.HostPort,
				},
			},
		},
		RestartPolicy: container.RestartPolicy{
			Name: params.PolicyName,
		},
	}
	id, err := dockerClient.ContainerCreate(containerOptions, hostOptions, nil, nil, params.ContainerName)
	if err != nil {
		response.Error(c, constant.ErrorCreateContainer, err, constant.ErrorMsg[constant.ErrorCreateContainer])
		return
	}
	response.OK(c, id, constant.ErrorMsg[constant.Success])
}

//func ContainerWait(c *gin.Context) {
//	var params common.ContainerLogsRequest
//	if err := c.ShouldBindJSON(&params); err != nil {
//		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
//		return
//	}
//
//	out, err := dockerClient.ContainerWait(params.ID)
//	if err != nil {
//		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
//		return
//	}
//	response.OK(c, out, constant.ErrorMsg[constant.Success])
//}
