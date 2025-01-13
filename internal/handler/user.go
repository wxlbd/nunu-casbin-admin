package handler

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wxlbd/nunu-casbin-admin/internal/dto"
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/service"
	"github.com/wxlbd/nunu-casbin-admin/pkg/config"
	"github.com/wxlbd/nunu-casbin-admin/pkg/ginx"
)

type UserHandler struct {
	svc service.Service
	cfg *config.Config
}

func NewUserHandler(svc service.Service, cfg *config.Config) *UserHandler {
	return &UserHandler{
		svc: svc,
		cfg: cfg,
	}
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c)
		return
	}

	accessToken, refreshToken, err := h.svc.User().Login(c, req.Username, req.Password)
	if err != nil {
		ginx.Unauthorized(c)
		return
	}

	ginx.Success(c, &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    int64(h.cfg.JWT.AccessExpire.Seconds()),
	})
}

// RefreshToken 刷新令牌
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c)
		return
	}

	accessToken, refreshToken, err := h.svc.User().RefreshToken(c, req.RefreshToken)
	if err != nil {
		ginx.Unauthorized(c)
		return
	}

	ginx.Success(c, &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    int64(h.cfg.JWT.AccessExpire.Seconds()),
	})
}

// Logout 用户登出
func (h *UserHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || len(token) <= 7 || token[:7] != "Bearer " {
		ginx.Error(c, 400, "无效的token")
		return
	}

	token = token[7:]
	if err := h.svc.User().Logout(c, token); err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}

	ginx.Success(c, nil)
}

// Create 创建用户
func (h *UserHandler) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ginx.Error(c, 400, "参数错误")
		return
	}

	if err := h.svc.User().Create(c, &user); err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}

	ginx.Success(c, nil)
}

// Update 更新用户
func (h *UserHandler) Update(c *gin.Context) {
	var user dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		ginx.Error(c, 400, "参数错误")
		return
	}

	if err := h.svc.User().Update(c, user.ToModel()); err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}

	ginx.Success(c, nil)
}

// Delete 删除用户
func (h *UserHandler) Delete(c *gin.Context) {
	str := strings.Split(c.Param("ids"), ",")
	ids := make([]uint64, 0, len(str))

	for _, s := range str {
		id, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			ginx.Error(c, 400, "参数错误")
			return
		}
		ids = append(ids, id)
	}
	if err := h.svc.User().Delete(c, ids...); err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}

	ginx.Success(c, nil)
}

// List 获取用户列表
func (h *UserHandler) List(c *gin.Context) {
	var req dto.UserListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ginx.Error(c, 400, "参数错误")
		return
	}

	// 参数验证
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 || req.PageSize > 100 {
		req.PageSize = 10
	}

	users, total, err := h.svc.User().List(c, req.ToModel())
	if err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}

	ginx.Success(c, dto.ToUserListResponse(users, total))
}

// AssignRoles 分配角色
func (h *UserHandler) AssignRoles(c *gin.Context) {
	var req struct {
		UserID  uint64   `json:"user_id" binding:"required"`
		RoleIDs []uint64 `json:"role_ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.Error(c, 400, "参数错误")
		return
	}

	if err := h.svc.User().AssignRoles(c, req.UserID, req.RoleIDs); err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}

	ginx.Success(c, nil)
}

// UpdatePassword 修改密码
func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var req dto.UpdatePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.Error(c, 400, "参数错误")
		return
	}

	if err := h.svc.User().UpdatePassword(c, req.ID, req.OldPassword, req.NewPassword); err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}

	ginx.Success(c, nil)
}

// Current 获取当前用户信息
func (h *UserHandler) Current(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		ginx.Error(c, 500, "获取用户信息失败")
		return
	}
	id, ok := userId.(uint64)
	if !ok {
		ginx.Error(c, 500, "获取用户信息失败")
		return
	}
	user, err := h.svc.User().FindByID(c.Request.Context(), id)
	if err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}
	ginx.Success(c, dto.ToUserResponse(user))
}

// Detail 获取当前用户信息
func (h *UserHandler) Detail(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		ginx.Error(c, 400, "参数错误")
		return
	}
	user, err := h.svc.User().FindByID(c.Request.Context(), id)
	if err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}
	ginx.Success(c, dto.ToUserResponse(user))
}

// GetRoles 获取用户角色列表
func (h *UserHandler) GetRoles(c *gin.Context) {
	id := c.GetUint64("user_id")
	if id == 0 {
		ginx.ParamError(c)
		return
	}

	// 获取用户的角色列表
	roles, err := h.svc.User().GetUserRoles(c, id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, roles)
}
