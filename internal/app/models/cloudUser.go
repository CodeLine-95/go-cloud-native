package models

import common "github.com/CodeLine-95/go-cloud-native/common/models"

type CloudUser struct {
	common.Model
	UserName  string `json:"user_name" gorm:"size:200;not null;uniqueIndex;default:'';comment:用户名"`
	PassWord  string `json:"pass_word" gorm:"size:255;not null;default:'';comment:密码"`
	UserEmail string `json:"user_email" gorm:"size:200;not null;default:'';comment:邮箱"`
	LoginIp   string `json:"login_ip" gorm:"size:200;not null;default:'';comment:登录IP"`
	LastTime  uint32 `json:"last_time" gorm:"index;not null;default:0;comment:最后登录时间"`
	Status    uint8  `json:"status" gorm:"index;not null;default:0;comment:用户状态: 0正常、1禁用"`
	Admin     uint8  `json:"admin" gorm:"index;not null;default:0;comment:是否超级管理员：0否、1是"`
	common.ModelTime
}

func (c CloudUser) TableName() string {
	return "cloud_user"
}
