package models

import (
	"encoding/json"
	common "github.com/CodeLine-95/go-cloud-native/common/models"
)

type CloudEtcd struct {
	common.Model
	Name       string `json:"name" gorm:"size:200;not null;uniqueIndex;default:'';comment:注册服务名"`
	Remark     string `json:"remark" gorm:"size:255;not null;index;default:'';comment:备注"`
	Content    string `json:"content" gorm:"text;not null;index;default:'';comment:注册内容"`
	IsSub      uint32 `json:"is_sub" gorm:"not null;index;default:0;comment:是否订阅"`
	SubUserID  string `json:"sub_user_id" gorm:"text;not null;index;default:'';comment:订阅的用户ID"`
	IsDelete   uint32 `json:"is_delete" gorm:"not null;index;default:0;comment:是否软删除"`
	IsRegister uint32 `json:"is_register" gorm:"not null;index;default:0;comment:是否注册成功"`
	common.ControlBy
	common.ModelTime
}

func (c CloudEtcd) TableName() string {
	return "cloud_etcd"
}

func (c *CloudEtcd) ParseFields(p any) *CloudEtcd {
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
