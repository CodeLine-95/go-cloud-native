package models

import (
	"encoding/json"
	common "github.com/CodeLine-95/go-cloud-native/common/models"
)

type CloudRoleMenu struct {
	RoleId uint32 `json:"role_id" gorm:"not null;index;default:0;comment:角色ID"`
	MenuId uint32 `json:"menu_id" gorm:"not null;index;default:0;comment:菜单ID"`
	common.ControlBy
	common.ModelTime
}

func (c CloudRoleMenu) TableName() string {
	return "cloud_role_menu"
}

func (c *CloudRoleMenu) ParseFields(p any) *CloudRoleMenu {
	if p == nil {
		return c
	}
	pjson, err := json.Marshal(p)
	if err != nil {
		return c
	}

	err = json.Unmarshal(pjson, c)
	if err != nil {
		return c
	}
	return c
}
