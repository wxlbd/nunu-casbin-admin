package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wxlbd/nunu-casbin-admin/internal/handler/request"
	"github.com/wxlbd/nunu-casbin-admin/internal/handler/response"
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/service"
)

type UserHandler struct {
	svc service.Service
}

func NewUserHandler(svc service.Service) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	accessToken, refreshToken, err := h.svc.User().Login(c, req.Username, req.Password)
	if err != nil {
		response.Unauthorized(c)
		return
	}

	response.Success(c, &response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// RefreshToken 刷新令牌
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var req request.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	accessToken, refreshToken, err := h.svc.User().RefreshToken(c, req.RefreshToken)
	if err != nil {
		response.Unauthorized(c)
		return
	}

	response.Success(c, &response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// Logout 用户登出
func (h *UserHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || len(token) <= 7 || token[:7] != "Bearer " {
		response.Error(c, 400, "无效的token")
		return
	}

	token = token[7:]
	if err := h.svc.User().Logout(c, token); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, nil)
}

// Create 创建用户
func (h *UserHandler) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, 400, "参数错误")
		return
	}

	if err := h.svc.User().Create(c, &user); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, nil)
}

// Update 更新用户
func (h *UserHandler) Update(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, 400, "参数错误")
		return
	}

	if err := h.svc.User().Update(c, &user); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, nil)
}

// Delete 删除用户
func (h *UserHandler) Delete(c *gin.Context) {
	var req struct {
		ID uint64 `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "参数错误")
		return
	}

	if err := h.svc.User().Delete(c, req.ID); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, nil)
}

// List 获取用户列表
func (h *UserHandler) List(c *gin.Context) {
	var req struct {
		Page int `form:"page" binding:"required,min=1"`
		Size int `form:"size" binding:"required,min=1,max=100"`
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, 400, "参数错误")
		return
	}

	users, total, err := h.svc.User().List(c, req.Page, req.Size)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  users,
		"total": total,
	})
}

// AssignRoles 分配角色
func (h *UserHandler) AssignRoles(c *gin.Context) {
	var req struct {
		UserID  uint64   `json:"user_id" binding:"required"`
		RoleIDs []uint64 `json:"role_ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "参数错误")
		return
	}

	if err := h.svc.User().AssignRoles(c, req.UserID, req.RoleIDs); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, nil)
}

// UpdatePassword 修改密码
func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var req struct {
		ID          uint64 `json:"id" binding:"required"`
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "参数错误")
		return
	}

	if err := h.svc.User().UpdatePassword(c, req.ID, req.OldPassword, req.NewPassword); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *UserHandler) Current(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		response.Error(c, 500, "获取用户信息失败")
		return
	}
	id, ok := userId.(uint64)
	if !ok {
		response.Error(c, 500, "获取用户信息失败")
		return
	}
	user, err := h.svc.User().FindByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	response.Success(c, user)
}
