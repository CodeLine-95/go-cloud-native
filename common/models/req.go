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
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Remark string `json:"remark,omitempty"`
	Key    string `json:"key,omitempty"`
	Sort   uint8  `json:"sort,omitempty"`
	Status uint8  `json:"status,omitempty"`
}

type ContainerStopRequest struct {
	Ids string `json:"ids"`
}

type ImagesPullRequest struct {
	Refstr string `json:"refstr"`
}

type MenuListRequest struct {
}

type MenuAddRequest struct {
}
type MenuEditRequest struct {
}
