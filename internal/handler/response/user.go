package response

// LoginResponse 登录响应
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

// UserResponse 用户信息响应
type UserResponse struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	UserType       int    `json:"user_type"`
	Nickname       string `json:"nickname"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Avatar         string `json:"avatar"`
	Signed         string `json:"signed"`
	Status         int    `json:"status"`
	LoginIp        string `json:"login_ip"`
	LoginTime      string `json:"login_time"`
	BackendSetting string `json:"backend_setting"`
	CreatedBy      int    `json:"created_by"`
	UpdatedBy      int    `json:"updated_by"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Remark         string `json:"remark"`
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	List  []*UserResponse `json:"list"`
	Total int64           `json:"total"`
}
