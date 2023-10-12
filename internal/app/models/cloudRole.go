package models

import common "github.com/CodeLine-95/go-cloud-native/common/models"

type CloudRole struct {
	RoleId     int64  `json:"role_id" xorm:"pk unsigned comment('唯一编号') version"`
	RoleName   string `json:"role_name" xorm:"varchar(200) notnull index comment('角色名')"`
	RoleKey    string `json:"role_key" xorm:"varchar(128) notnull index comment('角色Key')"`
	RoleRemark string `json:"role_remark" xorm:"varchar(255) notnull comment('角色备注')"`
	RoleSort   int64  `json:"role_sort" xorm:"tinyint(3) unsigned notnull index default(0) comment('角色排序')"`
	Status     int64  `json:"status" xorm:"tinyint(3) unsigned notnull index default(0) comment('角色状态: 0正常、1禁用')"`
	common.ControlBy
	common.ModelTime
}

func (c CloudRole) TableName() string {
	return "cloud_role"
}
