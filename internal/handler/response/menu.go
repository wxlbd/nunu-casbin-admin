package response

// 菜单树节点
type MenuTree struct {
	ID        uint64      `json:"id"`
	ParentID  uint64      `json:"parent_id"`
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	Component string      `json:"component"`
	Sort      int16       `json:"sort"`
	Status    int8        `json:"status"`
	Children  []*MenuTree `json:"children,omitempty"`
}

// 菜单列表响应
type MenuListResponse struct {
	List  []*MenuTree `json:"list"`
	Total int64       `json:"total"`
}

// 菜单信息响应
type Menu struct {
	ID         uint64 `json:"id"`
	ParentID   uint64 `json:"parent_id"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	Sort       int16  `json:"sort"`
	Status     int8   `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
