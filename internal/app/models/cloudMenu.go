package models

import common "github.com/CodeLine-95/go-cloud-native/common/models"

type CloudMenu struct {
	MenuId     int64  `json:"menu_id" xorm:"pk unsigned comment('唯一编号') version"`
	MenuName   string `json:"menu_name" xorm:"varchar(200) notnull comment('菜单名')"`
	MenuTitle  string `json:"menu_title" xorm:"varchar(255) notnull comment('菜单标题')"`
	MenuIcon   string `json:"menu_icon" xorm:"varchar(128) notnull comment('菜单图标')"`
	MenuPath   string `json:"menu_path" xorm:"varchar(255) notnull comment('菜单路径')"`
	PathGroup  string `json:"path_group" xorm:"varchar(255) notnull comment('菜单路径组')"`
	menuType   string `json:"menu_type" xorm:"varchar(255) notnull comment('菜单类型')"`
	MenuMethod string `json:"menu_method" xorm:"varchar(16) notnull comment('菜单请求类型')"`
	Permission string `json:"permission" xorm:"varchar(255) notnull comment('菜单权限标识')"`
	ParentId   int32  `json:"parent_id" xorm:"int(11) unsigned index notnull default(0) comment('父级ID')"`
	Component  string `json:"component" xorm:"varchar(255) notnull comment('菜单组件')"`
	MenuSort   int32  `json:"permission" xorm:"tinyint(1) unsigned index notnull default(0) comment('菜单排序')"`
	Visible    int32  `json:"visible" xorm:"tinyint(1) unsigned index notnull default(0) comment('菜单是否显示：0显示、1隐藏')"`
	IsFrame    int32  `json:"is_frame" xorm:"tinyint(1) unsigned index notnull default(0)"`
	common.ControlBy
	common.ModelTime
}

func (c CloudMenu) TableName() string {
	return "cloud_menu"
}
