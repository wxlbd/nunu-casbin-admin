package request

// 创建/更新角色请求
type RoleRequest struct {
	ID     uint64 `json:"id"` // 更新时必填
	Name   string `json:"name" binding:"required"`
	Code   string `json:"code" binding:"required"`
	Status int8   `json:"status"`
	Sort   int16  `json:"sort"`
	Remark string `json:"remark"`
}

// RoleIDRequest 删除角色请求
type RoleIDRequest struct {
	ID uint64 `json:"id" binding:"required"`
}

// 角色列表请求
type RoleListRequest struct {
	Page int `form:"page" binding:"required,min=1"`
	Size int `form:"size" binding:"required,min=1,max=100"`
}

// 分配菜单请求
type AssignMenusRequest struct {
	RoleID  uint64   `json:"role_id" binding:"required"`
	MenuIDs []uint64 `json:"menu_ids" binding:"required"`
}
