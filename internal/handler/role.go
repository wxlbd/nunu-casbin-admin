package handler

import (
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

// 更新角色
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

// 删除角色
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

// 获取角色列表
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

	// 转换响应结构
	var list []*response.RoleResponse
	for _, role := range roles {
		list = append(list, &response.RoleResponse{
			ID:     role.ID,
			Name:   role.Name,
			Code:   role.Code,
			Status: role.Status,
			Sort:   role.Sort,
			Remark: role.Remark,
		})
	}

	response.Success(c, &response.RoleListResponse{
		List:  list,
		Total: total,
	})
}

// 分配菜单
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
