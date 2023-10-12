package models

import common "github.com/CodeLine-95/go-cloud-native/common/models"

type CloudUser struct {
	common.Model
	UserName  string `json:"user_name" xorm:"varchar(255) notnull index comment('用户名')"`
	PassWord  string `json:"pass_word" xorm:"varchar(255) notnull comment('密码')"`
	UserEmail string `json:"user_email" xorm:"varchar(200) notnull comment('邮箱')"`
	LoginIp   string `json:"login_ip" xorm:"varchar(200) notnull comment('登录IP')"`
	LastTime  int64  `json:"last_time" xorm:"int(11) unsigned notnull comment('最后登录时间')"`
	Status    int64  `json:"status" xorm:"tinyint(3) unsigned notnull index default(0) comment('用户状态: 0正常、1禁用')"`
	Admin     int64  `json:"admin" xorm:"tinyint(3) unsigned notnull index default(0) comment('是否超级管理员：0否、1是')"`
	common.ModelTime
}

func (c CloudUser) TableName() string {
	return "cloud_user"
}
