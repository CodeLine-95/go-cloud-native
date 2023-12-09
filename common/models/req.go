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

type MenuRouterReqest struct {
	IsTree uint8 `json:"is_tree,omitempty"`
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
	NoCache    uint8  `json:"no_cache,omitempty"`
}

type UserRoleRequest struct {
	UID    uint32 `json:"uid"`
	RoleId uint32 `json:"role_id"`
}

type RoleMenuRequest struct {
	RoleId  uint32 `json:"role_id"`
	MenuIds string `json:"menu_ids"`
}

// 容器
type ContainerLogsRequest struct {
	ID string `json:"id"`
}
type BatchContainerLogsRequest struct {
	ID []string `json:"id"`
}
type ContainerStopRequest struct {
	Ids string `json:"ids"`
}
type ContainerCreateRequest struct {
	Image         string   `json:"image"`         // 指定要使用的镜像
	Cmd           []string `json:"cmd"`           // 指定容器启动时要执行的命令
	Hostname      string   `json:"hostname"`      // 主机名
	HostIP        string   `json:"hostIP"`        // 容器绑定IP
	LocalProt     string   `json:"localProt"`     // 容器绑定端口
	HostPort      string   `json:"hostPort"`      // 宿主机映射端口
	PolicyName    string   `json:"policyName"`    // 重启策略
	ContainerName string   `json:"containerName"` // 容器名称
	// 可选的重启策略：
	// - "no"：无重启策略
	// - "always"：容器总是自动重启
	// - "on-failure"：容器在非零退出状态时重启（默认最多重启3次）
	// - "unless-stopped"：除非手动停止，否则容器总是自动重启
}

type ImagesPullRequest struct {
	Refstr string `json:"refstr"`
}

type EtcdRequest struct {
	ID          int32  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Remark      string `json:"remark,omitempty"`
	Content     string `json:"content,omitempty"`
	IsSubscribe int32  `json:"is_subscribe,omitempty"`
	SubscribeID string `json:"subscribe_id,omitempty"`
}
