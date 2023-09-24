package store

import (
	"github.com/CodeLine-95/go-cloud-native/pkg/logz"
	"github.com/CodeLine-95/go-cloud-native/store/db"
)

func Init() {
	db.Init()
	logz.Info("db init succ")
}
