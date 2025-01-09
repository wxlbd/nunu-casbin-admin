package request

// 创建/更新菜单请求
type MenuRequest struct {
	ID        uint64 `json:"id"` // 更新时必填
	ParentID  uint64 `json:"parent_id"`
	Name      string `json:"name" binding:"required"`
	Path      string `json:"path" binding:"required"`
	Component string `json:"component" binding:"required"`
	Sort      int16  `json:"sort"`
	Status    int8   `json:"status"`
}

// 删除菜单请求
type MenuIDRequest struct {
	ID uint64 `json:"id" binding:"required"`
}

// 菜单列表请求
type MenuListRequest struct {
	Page int `form:"page" binding:"required,min=1"`
	Size int `form:"size" binding:"required,min=1,max=100"`
}
