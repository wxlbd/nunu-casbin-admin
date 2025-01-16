package dto

import (
	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/internal/types"
)

// UserBase 基础字段
type UserBase struct {
	Nickname       string                `json:"nickname"`
	Phone          string                `json:"phone"`
	Email          string                `json:"email"`
	Avatar         string                `json:"avatar"`
	Status         int8                  `json:"status"`
	UserType       int                   `json:"user_type"`
	Signed         string                `json:"signed"`
	BackendSetting *types.BackendSetting `json:"backend_setting"`
	Remark         string                `json:"remark"`
}

// CreateUserRequest CreateUserReq 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserBase
}

// UpdateUserRequest UpdateUserReq 更新用户请求
type UpdateUserRequest struct {
	ID       uint64 `json:"id" binding:"required"`
	Password string `json:"password,omitempty"` // 更新时密码可选
	UserBase
}

// UserResponse 用户响应
type UserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	UserBase
	LoginIp   string `json:"login_ip"`
	LoginTime string `json:"login_time"`
	CreatedBy uint64 `json:"created_by"`
	UpdatedBy uint64 `json:"updated_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ToModel 转换方法
func (req *CreateUserRequest) ToModel(createdBy uint64) *model.User {
	return &model.User{
		Username:  req.Username,
		Password:  req.Password,
		Nickname:  req.Nickname,
		Phone:     req.Phone,
		Email:     req.Email,
		Status:    req.Status,
		CreatedBy: createdBy,
	}
}

func (req *UpdateUserRequest) ToModel() *model.User {
	return &model.User{
		ID:       req.ID,
		Password: req.Password, // 如果为空则不会更新
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   req.Status,
		// UpdatedBy: updatedBy,
	}
}

// ToUserResponse 模型转响应
func ToUserResponse(m *model.User) *UserResponse {
	if m == nil {
		return nil
	}
	return &UserResponse{
		ID:       m.ID,
		Username: m.Username,
		UserBase: UserBase{
			Nickname:       m.Nickname,
			Phone:          m.Phone,
			Email:          m.Email,
			Status:         m.Status,
			BackendSetting: m.BackendSetting,
			Remark:         m.Remark,
			Avatar:         m.Avatar,
			UserType:       m.UserType,
			Signed:         m.Signed,
		},
		CreatedBy: m.CreatedBy,
		UpdatedBy: m.UpdatedBy,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
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

// UserListRequest 用户列表请求
type UserListRequest struct {
	*types.PageParam
	Username string `form:"username"`
	Nickname string `form:"nickname"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Status   int8   `form:"status"`
}

func (req *UserListRequest) ToModel() *model.UserQuery {
	return &model.UserQuery{
		Page:     req.Page,
		PageSize: req.PageSize,
		Username: req.Username,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   req.Status,
	}
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

// UserListResponse 用户列表响应
type UserListResponse struct {
	List  []*UserResponse `json:"list"`
	Total int64           `json:"total"`
}

// UserAssignRolesRequest 用户分配角色
type UserAssignRolesRequest struct {
	RoleCodes []string `json:"role_codes"`
}
