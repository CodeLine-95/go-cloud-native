package models

import common "github.com/CodeLine-95/go-cloud-native/common/models"

type CloudRole struct {
	RoleId     uint32 `json:"role_id" gorm:"primaryKey;autoIncrement;comment:唯一编号"`
	RoleName   string `json:"role_name" gorm:"size:200;not null;index;default:'';comment:角色名"`
	RoleKey    string `json:"role_key" gorm:"size:128;not null;index;default:'';comment:角色Key"`
	RoleRemark string `json:"role_remark" gorm:"size:255;not null;default:'';comment:角色备注"`
	RoleSort   uint8  `json:"role_sort" gorm:"not null;index;default:0;comment:角色排序"`
	Status     uint8  `json:"status" gorm:"not null;index;default:0;comment:角色状态: 0正常、1禁用"`
	common.ControlBy
	common.ModelTime
}

func (c CloudRole) TableName() string {
	return "cloud_role"
}
