package containers

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Docker struct {
	client.Client
}

func (d *Docker) GetClient() (cli *client.Client, err error) {
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return
}

func (d *Docker) ContainerList() (containerList []types.Container, err error) {
	cli, cliErr := d.GetClient()
	if cliErr != nil {
		panic(fmt.Sprintf("docker client error: %s", cliErr.Error()))
	}

	ctx := context.Background()

	containerOptions := types.ContainerListOptions{
		All: true,
	}

	containerList, err = cli.ContainerList(ctx, containerOptions)
	return
}
