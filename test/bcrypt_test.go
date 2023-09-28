package test

import (
	"flag"
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/pkg/logz"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/ip"
	"github.com/CodeLine-95/go-cloud-native/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/store"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
	"testing"
)

func parseConfig() {
	getwd, _ := os.Getwd()
	config := flag.String("c", filepath.Join(filepath.Dir(getwd), "/", "conf/local.toml"), "conf")
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

func TestBcrypt(t *testing.T) {
	//parseConfig()
	password := "123456"
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fromPassword))

	fmt.Println(ip.ClientIP())

	//engine := db.Grp("mysql")
	//
	//user := models.CloudUser{
	//	UserName:   "admin",
	//	PassWord:   string(fromPassword),
	//	UserEmail:  "admin@email.com",
	//	LoginIp:    ip.ClientIP(),
	//	CreateTime: int(time.Now().Unix()),
	//	Status:     1,
	//}
	//
	//cnt, err := engine.Insert(user)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(cnt)
}
