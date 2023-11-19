package docker

import (
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

// ImageList 获取镜像列表
func (d *DockerClient) ImageList() (imageList []types.ImageSummary, err error) {
	imageList, err = d.Client.ImageList(d.Ctx, types.ImageListOptions{
		All: true,
	})
	return
}

func (d *DockerClient) ImagePull(refStr string) {
	out, err := d.Client.ImagePull(d.Ctx, refStr, types.ImagePullOptions{
		All: true,
	})
	if err != nil {
		return
	}
	defer func(out io.ReadCloser) {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}(out)
	io.Copy(os.Stdout, out)
}
