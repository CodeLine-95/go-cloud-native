package models

import (
	"encoding/json"
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"gorm.io/gorm"
)

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

type GetCloudUserAndRole struct {
	CloudUser
	RoleId uint32 `json:"role_id"`
}

// AfterFind 查询记录后会调用它
func (cr *GetCloudUserAndRole) AfterFind(tx *gorm.DB) (err error) {
	if cr == nil {
		return
	}
	var cloudUserRole CloudUserRole
	ret := tx.Where("uid = ?", cr.Id).Find(&cloudUserRole)
	if ret.RowsAffected == 0 || ret.Error != nil {
		return
	}
	cr.RoleId = cloudUserRole.RoleId
	return
}

func (c *CloudUser) ParseFields(p any) *CloudUser {
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
