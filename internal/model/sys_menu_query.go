package model

type SysMenuQuery struct {
	Title    string `form:"title"`    // 菜单名称
	Status   int32  `form:"status"`   // 状态
	MenuType int32  `form:"type"`     // 菜单类型
	Page     int    `form:"page"`     // 页码
	PageSize int    `form:"pageSize"` // 每页数量
}

// SysMenuTree 菜单树结构
type SysMenuTree struct {
	*SysMenu
	Children []*SysMenuTree `json:"children"`
}
