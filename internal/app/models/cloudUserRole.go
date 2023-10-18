package models

import (
	"encoding/json"
	common "github.com/CodeLine-95/go-cloud-native/common/models"
)

type CloudUserRole struct {
	UID    uint32 `json:"uid" gorm:"primaryKey;autoIncrement;comment:用户ID"`
	RoleId uint32 `json:"role_id" gorm:"not null;index;default:0;comment:角色ID"`
	common.ControlBy
	common.ModelTime
}

func (c CloudUserRole) TableName() string {
	return "cloud_user_role"
}

func (c *CloudUserRole) ParseFields(p any) *CloudUserRole {
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
