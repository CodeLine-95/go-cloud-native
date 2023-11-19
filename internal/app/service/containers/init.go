package containers

import (
	"context"
	"github.com/CodeLine-95/go-cloud-native/internal/app/service/containers/docker"
	"github.com/docker/docker/client"
)

// 创建docker client 句柄
func InitDockerClient() *docker.DockerClient {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return &docker.DockerClient{
		Client: cli,
		Ctx:    context.Background(),
	}
}
