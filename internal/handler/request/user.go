package request

// 用户登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 刷新令牌请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// 创建/更新用户请求
type UserRequest struct {
	ID       uint64 `json:"id"` // 更新时必填
	Username string `json:"username" binding:"required"`
	Password string `json:"password,omitempty"` // 创建时必填，更新时选填
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   int8   `json:"status"`
}

// 删除用户请求
type UserIDRequest struct {
	ID uint64 `json:"id" binding:"required"`
}

// 用户列表请求
type UserListRequest struct {
	Page int `form:"page" binding:"required,min=1"`
	Size int `form:"size" binding:"required,min=1,max=100"`
}

// 分配角色请求
type AssignRolesRequest struct {
	UserID  uint64   `json:"user_id" binding:"required"`
	RoleIDs []uint64 `json:"role_ids" binding:"required"`
}

// 修改密码请求
type UpdatePasswordRequest struct {
	ID          uint64 `json:"id" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
