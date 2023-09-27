package server

import (
	"flag"
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/pkg/logz"
	"github.com/CodeLine-95/go-cloud-native/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/routers"
	"github.com/CodeLine-95/go-cloud-native/store"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "go-cloud-native",
	Short: "go-cloud-native start service",
}

func parseConfig() {
	config := flag.String("c", "conf/local.toml", "conf")
	flag.Parse()
	viper.SetConfigFile(*config)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("parse config file fail: %s", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logz.Info("Config: conf/local.toml Changed...")
	})

	// 初始化日志文件
	xlog.InitLog(viper.GetString("log.dir"), viper.GetString("log.level"), viper.GetString("log.name"))

	// 初始化数据库
	store.Init()
}

func init() {
	parseConfig()
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	gin.DisableConsoleColor()

	if viper.GetString("app.env") == "test" {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	// 关闭路由打印
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {}
	routers.InitRouter(r)

	port := viper.GetString("app.port")
	if port == "" {
		port = ":8080"
	}
	_ = r.Run(port)
}
