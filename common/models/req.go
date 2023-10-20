package models

type LoginRequest struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type SearchRequest struct {
	SearchKey string `json:"search_key,omitempty"`
	Page      int    `json:"page,omitempty"`
	PageSize  int    `json:"page_size,omitempty"`
}
type RoleRequest struct {
	RoleId     int    `json:"role_id,omitempty"`
	RoleName   string `json:"role_name,omitempty"`
	RoleRemark string `json:"role_remark,omitempty"`
	RoleKey    string `json:"role_key,omitempty"`
	RoleSort   uint8  `json:"role_sort,omitempty"`
	Status     uint8  `json:"status,omitempty"`
}

type MenuRequest struct {
	MenuId     uint32 `json:"menu_id,omitempty"`
	MenuName   string `json:"menu_name,omitempty"`
	MenuTitle  string `json:"menu_title,omitempty"`
	MenuIcon   string `json:"menu_icon,omitempty"`
	MenuPath   string `json:"menu_path,omitempty"`
	PathGroup  string `json:"path_group,omitempty"`
	MenuType   string `json:"menu_type,omitempty"`
	MenuMethod string `json:"menu_method,omitempty"`
	Permission string `json:"permission,omitempty"`
	ParentId   uint32 `json:"parent_id,omitempty"`
	Component  string `json:"component,omitempty"`
	MenuSort   uint8  `json:"menu_sort,omitempty"`
	Visible    uint8  `json:"visible,omitempty"`
	IsFrame    uint8  `json:"is_frame,omitempty"`
}

type UserRoleRequest struct {
	UID    uint32 `json:"uid"`
	RoleId uint32 `json:"role_id"`
}

type RoleMenuRequest struct {
	RoleId  uint32 `json:"role_id"`
	MenuIds string `json:"menu_ids"`
}

type ContainerStopRequest struct {
	Ids string `json:"ids"`
}

type ImagesPullRequest struct {
	Refstr string `json:"refstr"`
}
