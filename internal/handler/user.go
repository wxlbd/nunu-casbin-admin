package handler

import (
	"strconv"
	"strings"

	"github.com/wxlbd/gin-casbin-admin/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/wxlbd/gin-casbin-admin/internal/dto"
	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/ginx"
)

type UserHandler struct {
	svc Service
	cfg *config.Config
}

func NewUserHandler(svc Service, cfg *config.Config) *UserHandler {
	return &UserHandler{
		svc: svc,
		cfg: cfg,
	}
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录并获取访问令牌
// @Tags 认证管理
// @Accept json
// @Produce json
// @Param data body dto.LoginRequest true "登录信息"
// @Success 200 {object} ginx.Response{data=dto.LoginResponse} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 401 {object} ginx.Response "用户名或密码错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	accessToken, refreshToken, err := h.svc.User().Login(c, req.Username, req.Password)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    int64(h.cfg.JWT.AccessExpire.Seconds()),
	})
}

// RefreshToken 刷新令牌
// @Summary 刷新访问令牌
// @Description 使用刷新令牌获取新的访问令牌
// @Tags 认证管理
// @Accept json
// @Produce json
// @Param data body dto.RefreshTokenRequest true "刷新令牌"
// @Success 200 {object} ginx.Response{data=dto.LoginResponse} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 401 {object} ginx.Response "令牌无效或已过期"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Router /auth/refresh-token [post]
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
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
	// 未登录或非法访问直接返回
	if token == "" || len(token) <= 7 || token[:7] != "Bearer " {
		ginx.Success(c, nil)
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
// @Summary 创建用户
// @Description 创建一个新的用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param data body dto.CreateUserRequest true "用户信息"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/user [post]
func (h *UserHandler) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ginx.ParamError(c, err)
		return
	}

	if err := h.svc.User().Create(c, &user); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Update 更新用户
// @Summary 更新用户
// @Description 更新指定ID的用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param data body dto.UpdateUserRequest true "用户信息"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "用户不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/user/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	var user dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		ginx.ParamError(c, err)
		return
	}

	if err := h.svc.User().Update(c, user.ToModel()); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Delete 删除用户
// @Summary 删除用户
// @Description 删除指定ID的用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param ids path string true "用户ID列表(多个用逗号分隔)"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/user/{ids} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	str := strings.Split(c.Param("ids"), ",")
	ids := make([]uint64, 0, len(str))

	for _, s := range str {
		id, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
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
// @Summary 获取用户列表
// @Description 分页获取用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param username query string false "用户名"
// @Param nickname query string false "昵称"
// @Param status query int false "状态(1:正常 2:禁用)"
// @Success 200 {object} ginx.Response{data=ginx.ListData{list=[]dto.UserResponse,total=int64}} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/user [get]
func (h *UserHandler) List(c *gin.Context) {
	var req dto.UserListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	req.Normalize()
	users, total, err := h.svc.User().List(c, req.ToModel())
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, dto.ToUserListResponse(users, total))
}

// AssignRoles 分配角色
// @Summary 分配角色
// @Description 为指定用户分配角色
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param data body dto.UserAssignRolesRequest true "角色列表"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "用户不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/user/{id}/roles [post]
func (h *UserHandler) AssignRoles(c *gin.Context) {
	var req dto.UserAssignRolesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
		return
	}
	if err := h.svc.User().AssignRoles(c, userID, req.RoleCodes); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// UpdatePassword 修改密码
func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var req dto.UpdatePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	if err := h.svc.User().UpdatePassword(c, req.ID, req.OldPassword, req.NewPassword); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Current 获取当前用户信息
// @Summary 获取当前用户信息
// @Description 获取当前登录用户的详细信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {object} ginx.Response{data=dto.UserResponse} "成功"
// @Failure 401 {object} ginx.Response "未登录或令牌无效"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /profile [get]
func (h *UserHandler) Current(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
		return
	}
	id, ok := userId.(uint64)
	if !ok {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
		return
	}
	user, err := h.svc.User().FindByID(c.Request.Context(), id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, dto.ToUserResponse(user))
}

// Detail 获取当前用户信息
func (h *UserHandler) Detail(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
		return
	}
	user, err := h.svc.User().FindByID(c.Request.Context(), id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, dto.ToUserResponse(user))
}

// GetCurrentUserRoles 获取当前用户角色列表
func (h *UserHandler) GetCurrentUserRoles(c *gin.Context) {
	id := c.GetUint64("user_id")
	if id == 0 {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
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

func (h *UserHandler) GerUserRoles(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(ctx, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
		return
	}
	roles, err := h.svc.User().GetUserRoles(ctx.Request.Context(), id)
	if err != nil {
		ginx.ServerError(ctx, err)
		return
	}
	ginx.Success(ctx, roles)
}
