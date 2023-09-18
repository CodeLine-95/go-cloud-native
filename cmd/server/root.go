package server

import (
	"fmt"
	routersR "github.com/CodeLine-95/go-cloud-native/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "go-cloud-native",
	Short: "go-cloud-native start service",
}

func init() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	r := gin.Default()
	fmt.Println(r)
	r = routersR.ApiV1(r)
	_ = r.Run(":8000")
}
