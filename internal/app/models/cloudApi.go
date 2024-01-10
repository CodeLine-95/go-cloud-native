package models

import (
	"encoding/json"
	"github.com/CodeLine-95/go-cloud-native/common/models"
)

type CloudApi struct {
	models.Model
	ID         uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:唯一编号"`
	ApiName    string `json:"api_name" gorm:"size:200;not null;default:'';comment:接口名称"`
	ApiKey     string `json:"api_key" gorm:"size:100;not null;uniqueIndex;default:'';comment:接口标识"`
	ApiMethod  string `json:"api_method" gorm:"size:50;not null;default:'';comment:接口请求类型：POST、GET、PUT、DELETE"`
	ApiDesc    string `json:"api_desc" gorm:"size:128;not null;default:'';comment:接口描述"`
	ApiUrl     string `json:"api_url" gorm:"size:255;not null;uniqueIndex;default:'';comment:接口地址"`
	CreateUser string `json:"create_user" gorm:"size:255;not null;default:'';comment:接口创建人"`
	UpdateUser string `json:"update_user" gorm:"size:255;not null;default:'';comment:接口修改人"`
	ApiSort    uint8  `json:"api_sort" gorm:"index;not null;default:0;comment:接口排序"`
	ApiStatus  uint8  `json:"api_status" gorm:"index;not null;default:0;comment:接口状态（0未完成，1已完成，2已废弃）"`
	models.ModelTime
}

func (c CloudApi) TableName() string {
	return "cloud_api"
}

// ParseFields 提取tag值
func (c *CloudApi) ParseFields(p any) *CloudApi {
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
