package initial

import (
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/initial/store/web"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/spf13/viper"
)

func Init() {
	// 初始化日志文件
	xlog.InitLog(
		viper.GetString("log.dir"),
		viper.GetString("log.level"),
		viper.GetString("log.name"),
	)

	// 初始化DB
	db.Init()

	// 初始化 web 框架
	web.Init()
}
