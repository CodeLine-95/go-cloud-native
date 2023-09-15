package main

import (
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/cmd/server"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	if err := server.RootCmd.Execute(); err != nil {
		panic(fmt.Errorf("cmd exec fail: %s", err))
	}
}
