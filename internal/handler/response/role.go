package response

// RoleResponse 角色信息响应
type RoleResponse struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Status  int8   `json:"status"`
	Sort    int16  `json:"sort"`
	Remark  string `json:"remark"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

// RoleListResponse 角色列表响应
type RoleListResponse struct {
	List  []*RoleResponse `json:"list"`
	Total int64           `json:"total"`
}

// RoleMenusResponse 角色菜单响应
type RoleMenusResponse []*Menu
