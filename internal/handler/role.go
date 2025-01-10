package handler

import (
	"github.com/wxlbd/nunu-casbin-admin/internal/dto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wxlbd/nunu-casbin-admin/internal/handler/request"
	"github.com/wxlbd/nunu-casbin-admin/internal/handler/response"
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/service"
)

type RoleHandler struct {
	svc service.Service
}

func NewRoleHandler(svc service.Service) *RoleHandler {
	return &RoleHandler{
		svc: svc,
	}
}

// Create 创建角色
func (h *RoleHandler) Create(c *gin.Context) {
	var req request.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	role := &model.Role{
		Name:   req.Name,
		Code:   req.Code,
		Status: req.Status,
		Sort:   req.Sort,
		Remark: req.Remark,
	}

	if err := h.svc.Role().Create(c, role); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}

// Update 更新角色
func (h *RoleHandler) Update(c *gin.Context) {
	var req request.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	role := &model.Role{
		ID:     req.ID,
		Name:   req.Name,
		Code:   req.Code,
		Status: req.Status,
		Sort:   req.Sort,
		Remark: req.Remark,
	}

	if err := h.svc.Role().Update(c, role); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}

// Delete 删除角色
func (h *RoleHandler) Delete(c *gin.Context) {
	var req request.RoleIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	if err := h.svc.Role().Delete(c, req.ID); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}

// List 获取角色列表
func (h *RoleHandler) List(c *gin.Context) {
	var req request.RoleListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ParamError(c)
		return
	}

	roles, total, err := h.svc.Role().List(c, req.Page, req.Size)
	if err != nil {
		response.ServerError(c, err)
		return
	}
	response.Success(c, &response.RoleListResponse{
		List:  dto.ToRoleList(roles),
		Total: total,
	})
}

// AssignMenus 分配菜单
func (h *RoleHandler) AssignMenus(c *gin.Context) {
	var req request.AssignMenusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	if err := h.svc.Role().AssignMenus(c, req.RoleID, req.MenuIDs); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}

// GetMenus 获取角色菜单
func (h *RoleHandler) GetMenus(c *gin.Context) {
	// 从查询参数获取角色ID
	roleID := c.Query("role_id")
	if roleID == "" {
		response.ParamError(c)
		return
	}

	// 转换为 uint64
	id, err := strconv.ParseUint(roleID, 10, 64)
	if err != nil {
		response.ParamError(c)
		return
	}

	// 获取角色的菜单列表
	menus, err := h.svc.Role().GetRoleMenus(c, id)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, menus)
}

func (h *RoleHandler) Detail(ctx *gin.Context) {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}
	role, err := h.svc.Role().FindByID(ctx.Request.Context(), id)
	if err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, dto.ToRoleResponse(role))
}
