package server

import (
	"flag"
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/routers"
	logz "github.com/CodeLine-95/go-tools"
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
}

func init() {
	parseConfig()
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	if viper.GetString("app.env") == "test" {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()
	r = routers.ApiV1(r)

	port := viper.GetString("app.port")
	if port == "" {
		port = ":8080"
	}
	_ = r.Run(port)
}
