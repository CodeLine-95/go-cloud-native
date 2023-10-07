package main

import (
	"flag"
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/initial"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func parseConfig() {
	getwd, _ := os.Getwd()
	config := flag.String("c", filepath.Join(getwd, "/", "configs/local.toml"), "conf")
	flag.Parse()
	viper.SetConfigFile(*config)
	if err := viper.ReadInConfig(); err != nil {
		logz.Error("parse config file fail", logz.F("err", err))
		panic(fmt.Sprintf("parse config file fail: %s", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logz.Info("Config: configs/local.toml Changed...")
	})
	// 初始化配置
	initial.Init()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			logz.Error("server start fail...", logz.F("err", err))
			panic(err)
		}
	}()
	parseConfig()
}
