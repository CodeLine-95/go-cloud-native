package routers

import (
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/middleware"
	routers "github.com/CodeLine-95/go-cloud-native/routers/docker"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var handlersFuncMap []gin.HandlerFunc

func init() {
	handlersFuncMap = append(handlersFuncMap, middleware.JWTLogin())
}

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Logger(), middleware.Recovery())
	versionRouter := r.Group(fmt.Sprintf("/%s", viper.GetString("app.apiVersion")))
	Common(versionRouter)

	// 批量设置中间件:  jwt登录验证
	versionRouter.Use(handlersFuncMap...)
	routers.DockerRouter(versionRouter)
	return r
}
