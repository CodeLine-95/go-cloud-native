package web

import (
	"github.com/CodeLine-95/go-cloud-native/initial/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func Init() {
	// 终端输出彩色日志
	gin.DisableConsoleColor()

	//初始化 gin
	r := gin.New()

	// 测试环境打开 debug
	if viper.GetString("app.env") == "test" || viper.GetString("app.env") == "dev" {
		gin.SetMode(gin.DebugMode)
		r.Use(gin.Logger())
	}

	// Debug 日志输出格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("[cloud-native] %v %v %v \n", httpMethod, absolutePath, handlerName)
	}

	// 加载路由配置
	router.InitRouter(r)

	// 监听服务端口
	port := viper.GetString("app.port")
	if port == "" {
		port = ":8080"
	}
	_ = r.Run(port)
}
