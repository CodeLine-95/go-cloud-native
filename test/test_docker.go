package main

import (
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/pkg/containers"
)

func main() {
	docker := containers.Docker{}

	containerList, err := docker.ContainerList()
	if err != nil {
		panic(err)
	}

	fmt.Println(containerList)
}
