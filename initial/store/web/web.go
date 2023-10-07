package web

import (
	"github.com/CodeLine-95/go-cloud-native/initial/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	// 终端输出彩色日志
	gin.DisableConsoleColor()

	// 测试环境打开 debug
	if viper.GetString("app.env") == "test" {
		gin.SetMode(gin.DebugMode)
	}

	//初始化 gin
	r := gin.New()

	// 关闭路由打印
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {}

	// 加载路由配置
	router.InitRouter(r)

	// 监听服务端口
	port := viper.GetString("app.port")
	if port == "" {
		port = ":8080"
	}
	_ = r.Run(port)
}
