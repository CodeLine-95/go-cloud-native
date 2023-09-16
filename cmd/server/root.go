package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "go-cloud-native",
	Short: "go-cloud-native start service",
}

func init() {
	r := gin.Default()
	fmt.Println(r)
}
