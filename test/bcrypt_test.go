package test

import (
	"flag"
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/initial"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
	"testing"
	"time"
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
	// 初始化数据库
	initial.Init()
}

func TestBcrypt(t *testing.T) {
	parseConfig()
	password := "123456"
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fromPassword))

	//fmt.Println(ip.ClientIP())

	engine := db.Grp(constant.CloudNative)

	user := models.CloudUser{
		UserName:   "admin",
		PassWord:   string(fromPassword),
		UserEmail:  "admin@email.com",
		CreateTime: time.Now().Unix(),
		Status:     1,
	}

	cnt, err := engine.Insert(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(cnt)
}
