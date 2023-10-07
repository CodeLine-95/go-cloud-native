package containers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Docker struct {
	client.Client
}

func (d *Docker) GetClient() (cli *client.Client, err error) {
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return
}

// ContainerList 获取容器列表
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

// ContainerStop 停止容器
func (d *Docker) ContainerStop(containerID []string) (code int, err error) {
	cli, cliErr := d.GetClient()
	if cliErr != nil {
		panic(fmt.Sprintf("docker client error: %s", cliErr.Error()))
	}
	ctx := context.Background()

	c := new(gin.Context)

	for _, container := range containerID {
		err := cli.ContainerStop(ctx, container, nil)
		if err != nil {
			xlog.Error(traceId.GetLogContext(c, "container stop fail",
				logz.F("err", err),
			))
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}

// ContainerLogs 获取指定容器日志
func (d *Docker) ContainerLogs(container string) string {
	cli, cliErr := d.GetClient()
	if cliErr != nil {
		panic(fmt.Sprintf("docker client error: %s", cliErr.Error()))
	}
	ctx := context.Background()
	c := new(gin.Context)

	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, container, options)
	if err != nil {
		xlog.Error(traceId.GetLogContext(c, "container log ID fail",
			logz.F("err", err),
		))
		return ""
	}
	defer func(out io.ReadCloser) {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}(out)
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	return buf.String()
}

// ImageList 获取镜像列表
func (d *Docker) ImageList() (containerList []types.ImageSummary, err error) {
	cli, cliErr := d.GetClient()
	if cliErr != nil {
		panic(fmt.Sprintf("docker client error: %s", cliErr.Error()))
	}

	ctx := context.Background()

	imageListOptions := types.ImageListOptions{
		All: true,
	}

	containerList, err = cli.ImageList(ctx, imageListOptions)
	return
}

// ImagePull 拉取指定镜像
func (d *Docker) ImagePull(refstr string) string {
	cli, cliErr := d.GetClient()
	if cliErr != nil {
		panic(fmt.Sprintf("docker client error: %s", cliErr.Error()))
	}
	ctx := context.Background()
	c := new(gin.Context)

	options := types.ImagePullOptions{All: true}
	out, err := cli.ImagePull(ctx, refstr, options)
	if err != nil {
		xlog.Error(traceId.GetLogContext(c, fmt.Sprintf("images pull %s fail", refstr),
			logz.F("err", err),
		))
		return ""
	}
	defer func(out io.ReadCloser) {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}(out)
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	return buf.String()
}
