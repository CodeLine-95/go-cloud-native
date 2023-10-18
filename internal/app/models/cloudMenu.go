package models

import (
	"encoding/json"
	common "github.com/CodeLine-95/go-cloud-native/common/models"
)

type CloudMenu struct {
	MenuId     uint32 `json:"menu_id" gorm:"primaryKey;autoIncrement;comment:唯一编号"`
	MenuName   string `json:"menu_name" gorm:"size:200;not null;default:'';comment:菜单名"`
	MenuTitle  string `json:"menu_title" gorm:"size:128;not null;default:'';comment:菜单标题"`
	MenuIcon   string `json:"menu_icon" gorm:"size:128;not null;default:'';comment:菜单图标"`
	MenuPath   string `json:"menu_path" gorm:"size:255;not null;uniqueIndex;default:'';comment:菜单路径"`
	PathGroup  string `json:"path_group" gorm:"size:255;not null;default:'';comment:菜单路径组"`
	MenuType   string `json:"menu_type" gorm:"size:255;not null;default:'';comment:菜单类型"`
	MenuMethod string `json:"menu_method" gorm:"size:16;not null;default:'';comment:菜单请求类型"`
	Permission string `json:"permission" gorm:"size:255;not null;uniqueIndex;default:'';comment:菜单权限标识"`
	ParentId   uint32 `json:"parent_id" gorm:"index;not null;default:0;comment:父级ID"`
	Component  string `json:"component" gorm:"size:255;not null;uniqueIndex;default:'';comment:菜单组件"`
	MenuSort   uint8  `json:"menu_sort" gorm:"index;not null;default:0;comment:菜单排序"`
	Visible    uint8  `json:"visible" gorm:"index;not null;default:0;comment:菜单是否显示：0显示、1隐藏"`
	IsFrame    uint8  `json:"is_frame" gorm:"index;not null;default:0"`
	common.ControlBy
	common.ModelTime
	// 树节点
	ChildNode *CloudMenuTree `json:"child_node" gorm:"-"`
}

func (c CloudMenu) TableName() string {
	return "cloud_menu"
}

// ParseFields 提取tag值
func (c *CloudMenu) ParseFields(p any) *CloudMenu {
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

type CloudMenuTree []*CloudMenu

// TreeNode 格式化树节点
func (c CloudMenuTree) TreeNode() CloudMenuTree {
	if len(c) <= 0 {
		return c
	}
	// 先重组数据：以数据的ID作为外层的key编号，以便下面进行子树的数据组合
	TreeMenuData := make(map[uint32]*CloudMenu)
	for _, item := range c {
		TreeMenuData[item.MenuId] = item
	}
	var TreeNode CloudMenuTree
	for _, val := range TreeMenuData {
		if val.ParentId == 0 {
			TreeNode = append(TreeNode, val)
			continue
		}
		if p_item, ok := TreeMenuData[val.ParentId]; ok {
			if p_item.ChildNode == nil {
				p_item.ChildNode = &CloudMenuTree{val}
				continue
			}
			*p_item.ChildNode = append(*p_item.ChildNode, val)
		}
	}
	return TreeNode
}
