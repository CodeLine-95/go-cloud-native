package models

type CloudUser struct {
	Id         int64  `json:"id" xorm:"pk comment('唯一编号') version"`
	UserName   string `json:"user_name" xorm:"varchar(255) notnull index comment('用户名')"`
	PassWord   string `json:"pass_word" xorm:"varchar(255) notnull comment('密码')"`
	UserEmail  string `json:"user_email" xorm:"varchar(200) notnull comment('邮箱')"`
	LoginIp    string `json:"login_ip" xorm:"varchar(200) notnull comment('登录IP')"`
	LastTime   int64  `json:"last_time" xorm:"int(11) notnull comment('最后登录时间')"`
	CreateTime int64  `json:"create_time" xorm:"int(11) comment('创建时间')"`
	Status     int64  `json:"status" xorm:"int(10) notnull index default(0) comment('用户状态: 0正常、1禁用')"`
}

func (c CloudUser) TableName() string {
	return "cloud_user"
}
