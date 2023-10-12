package models

import common "github.com/CodeLine-95/go-cloud-native/common/models"

type CloudRoleRule struct {
	RoleId int64 `json:"role_id" xorm:"int(11) unsigned notnull index default(0) comment('角色ID')"`
	MenuId int64 `json:"menu_id" xorm:"int(11) unsigned notnull index default(0) comment('菜单ID')"`
	common.ModelTime
}

func (c CloudRoleRule) TableName() string {
	return "cloud_role_rule"
}
