package dto

import (
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/types"
	"time"
)

// ToUserResponse 将 User 模型转换为响应 DTO
func ToUserResponse(user *model.User) *UserResponse {
	if user == nil {
		return nil
	}
	return &UserResponse{
		ID:             int(user.ID),
		Username:       user.Username,
		Nickname:       user.Nickname,
		Phone:          user.Phone,
		Email:          user.Email,
		Avatar:         user.Avatar,
		Status:         int(user.Status),
		LoginTime:      user.LoginTime.Format(time.DateTime),
		CreatedBy:      int(user.CreatedBy),
		UpdatedBy:      int(user.UpdatedBy),
		CreatedAt:      user.CreatedAt.Format(time.DateTime),
		UpdatedAt:      user.UpdatedAt.Format(time.DateTime),
		Remark:         user.Remark,
		UserType:       user.UserType,
		Signed:         user.Signed,
		LoginIp:        user.LoginIp,
		BackendSetting: user.BackendSetting,
	}
}

// ToUserResponseList 将用户列表转换为响应 DTO 列表
func ToUserResponseList(users []*model.User) []*UserResponse {
	if users == nil {
		return nil
	}
	list := make([]*UserResponse, 0, len(users))
	for _, user := range users {
		list = append(list, ToUserResponse(user))
	}
	return list
}

// ToUserListResponse 转换为分页响应
func ToUserListResponse(users []*model.User, total int64) *UserListResponse {
	return &UserListResponse{
		List:  ToUserResponseList(users),
		Total: total,
	}
}

// LoginRequest 用户登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RefreshTokenRequest 刷新令牌请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UserRequest 创建/更新用户请求
type UserRequest struct {
	ID             uint64                `json:"id"` // 更新时必填
	Username       string                `json:"username" binding:"required"`
	Password       string                `json:"password,omitempty"` // 创建时必填，更新时选填
	Nickname       string                `json:"nickname"`
	Phone          string                `json:"phone"`
	Email          string                `json:"email"`
	Status         int8                  `json:"status"`
	BackendSetting *types.BackendSetting `json:"backend_setting"`
	UserType       int                   `json:"user_type"`
	Remark         string                `json:"remark"`
}

// UserListRequest 用户列表请求
type UserListRequest struct {
	types.PageParam
}

// AssignRolesRequest 分配角色请求
type AssignRolesRequest struct {
	UserID  uint64   `json:"user_id" binding:"required"`
	RoleIDs []uint64 `json:"role_ids" binding:"required"`
}

// UpdatePasswordRequest 修改密码请求
type UpdatePasswordRequest struct {
	ID          uint64 `json:"id" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

// UserResponse 用户信息响应
type UserResponse struct {
	ID             int                   `json:"id"`
	Username       string                `json:"username"`
	UserType       int                   `json:"user_type"`
	Nickname       string                `json:"nickname"`
	Phone          string                `json:"phone"`
	Email          string                `json:"email"`
	Avatar         string                `json:"avatar"`
	Signed         string                `json:"signed"`
	Status         int                   `json:"status"`
	LoginIp        string                `json:"login_ip"`
	LoginTime      string                `json:"login_time"`
	BackendSetting *types.BackendSetting `json:"backend_setting"`
	CreatedBy      int                   `json:"created_by"`
	UpdatedBy      int                   `json:"updated_by"`
	CreatedAt      string                `json:"created_at"`
	UpdatedAt      string                `json:"updated_at"`
	Remark         string                `json:"remark"`
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	List  []*UserResponse `json:"list"`
	Total int64           `json:"total"`
}
