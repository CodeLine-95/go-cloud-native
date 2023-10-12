package models

type LoginRequest struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type RoleListRequest struct {
	SearchKey string `json:"search_key"`
	Page      int64  `json:"page"`
	PageSize  int64  `json:"page_size"`
}

type RoleAddRequest struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Key    string `json:"key"`
}
type RoleEditRequest struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Key    string `json:"key"`
	Sort   int64  `json:"sort"`
	Status int64  `json:"status"`
}

type ContainerStopRequest struct {
	Ids string `json:"ids"`
}

type ImagesPullRequest struct {
	Refstr string `json:"refstr"`
}
