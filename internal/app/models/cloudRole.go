package models

type CloudRole struct {
	Id         int64  `json:"id" xorm:"pk comment('唯一编号') version"`
	Name       string `json:"name" xorm:"varchar(255) notnull index comment('角色名')"`
	Remark     string `json:"remark" xorm:"varchar(255) notnull comment('角色备注')"`
	RulesIds   string `json:"rules_ids" xorm:"MEDIUMTEXT notnull comment('权限ID')"`
	Status     int64  `json:"status" xorm:"int(10) notnull index default(0) comment('角色状态: 0正常、1禁用')"`
	CreatTime  int64  `json:"creat_time" xorm:"int(11) comment('创建时间')"`
	UpdateTime int64  `json:"update_time" xorm:"int(11) comment('更新时间')"`
}

func (c CloudRole) TableName() string {
	return "cloud_role"
}
