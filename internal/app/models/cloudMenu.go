package models

import common "github.com/CodeLine-95/go-cloud-native/common/models"

type CloudMenu struct {
	MenuId     uint32 `json:"menu_id" gorm:"primaryKey;autoIncrement;comment:唯一编号"`
	MenuName   string `json:"menu_name" gorm:"size:200;not null;default:'';comment:菜单名"`
	MenuTitle  string `json:"menu_title" gorm:"size:128;not null;default:'';comment:菜单标题"`
	MenuIcon   string `json:"menu_icon" gorm:"size:128;not null;default:'';comment:菜单图标"`
	MenuPath   string `json:"menu_path" gorm:"size:255;not null;uniqueIndex;default:'';comment:菜单路径"`
	PathGroup  string `json:"path_group" gorm:"size:255;not null;default:'';comment:菜单路径组"`
	menuType   string `json:"menu_type" gorm:"size:255;not null;default:'';comment:菜单类型"`
	MenuMethod string `json:"menu_method" gorm:"size:16;not null;default:'';comment:菜单请求类型"`
	Permission string `json:"permission" gorm:"size:255;not null;uniqueIndex;default:'';comment:菜单权限标识"`
	ParentId   uint32 `json:"parent_id" gorm:"index;not null;default:0;comment:父级ID"`
	Component  string `json:"component" gorm:"size:255;not null;uniqueIndex;default:'';comment:菜单组件"`
	MenuSort   uint8  `json:"permission" gorm:"index;not null;default:0;comment:菜单排序"`
	Visible    uint8  `json:"visible" gorm:"index;not null;default:0;comment:菜单是否显示：0显示、1隐藏"`
	IsFrame    uint8  `json:"is_frame" gorm:"index;not null;default:0"`
	common.ControlBy
	common.ModelTime
}

func (c CloudMenu) TableName() string {
	return "cloud_menu"
}
