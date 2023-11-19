package docker

import (
	"context"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	Client *client.Client
	Ctx    context.Context
}
