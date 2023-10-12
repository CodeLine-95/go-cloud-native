package test

import (
	"flag"
	"fmt"
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
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
	config := flag.String("c", filepath.Join(filepath.Dir(getwd), "/", "configs/local.toml"), "conf")
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
	db.Init()
}

func TestBcrypt(t *testing.T) {
	parseConfig()
	password := "123456"
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fromPassword))

	user := &models.CloudUser{
		UserName:  "admin",
		PassWord:  string(fromPassword),
		UserEmail: "admin@email.com",
		Status:    1,
		ModelTime: common.ModelTime{CreateTime: uint32(time.Now().Unix())},
	}

	res := db.D().Create(user)
	if res.Error != nil {
		panic(err)
	}

	fmt.Println("insert row count: ", res.RowsAffected)
}
