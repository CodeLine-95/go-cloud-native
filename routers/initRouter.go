package routers

import (
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/middleware"
	routers "github.com/CodeLine-95/go-cloud-native/routers/docker"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Logger(), middleware.Recovery())
	versionRouter := r.Group(fmt.Sprintf("/%s", viper.GetString("app.version")))
	routers.DockerRouter(versionRouter)
	return r
}
