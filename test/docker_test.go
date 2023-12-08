package test

import (
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/internal/app/service/containers"
	"testing"
)

func TestDocker(t *testing.T) {
	dockerClient := containers.NewDockerClient()

	containerList, err := dockerClient.ContainerList()
	if err != nil {
		panic(err)
	}

	fmt.Println(containerList)
}
