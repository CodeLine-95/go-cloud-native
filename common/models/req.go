package models

type LoginRequest struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type RoleListRequest struct {
	SearchKey string `json:"search_key"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}

type RoleAddRequest struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Key    string `json:"key"`
	Sort   uint8  `json:"sort"`
	Status uint8  `json:"status"`
}
type RoleEditRequest struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Key    string `json:"key"`
	Sort   uint8  `json:"sort"`
	Status uint8  `json:"status"`
}

type ContainerStopRequest struct {
	Ids string `json:"ids"`
}

type ImagesPullRequest struct {
	Refstr string `json:"refstr"`
}
