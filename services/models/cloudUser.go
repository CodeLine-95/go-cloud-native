package models

import "github.com/CodeLine-95/go-cloud-native/store/db"

type CloudUser struct {
	ID         int    `json:"id"`
	UserName   string `json:"user_name"`
	PassWord   string `json:"pass_word"`
	UserEmail  string `json:"user_email"`
	LoginIp    string `json:"login_ip"`
	LastTime   int    `json:"last_time"`
	CreateTime int    `json:"create_time"`
	Status     int    `json:"status"`
}

func (c CloudUser) TableName() string {
	tableName := db.GetTableName("cloud_user")
	if tableName != "" {
		return tableName
	}
	return "cloud_user"
}
