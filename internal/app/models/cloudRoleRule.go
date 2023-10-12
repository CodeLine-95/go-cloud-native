package models

import common "github.com/CodeLine-95/go-cloud-native/common/models"

type CloudRoleRule struct {
	RoleId uint32 `json:"role_id" gorm:"not null;index;default:0;comment:角色ID"`
	MenuId uint32 `json:"menu_id" gorm:"not null;index;default:0;comment:菜单ID"`
	common.ModelTime
}

func (c CloudRoleRule) TableName() string {
	return "cloud_role_rule"
}
