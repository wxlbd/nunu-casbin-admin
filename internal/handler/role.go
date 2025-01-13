package handler

import (
	"strconv"
	"strings"

	"github.com/wxlbd/nunu-casbin-admin/internal/dto"
	"github.com/wxlbd/nunu-casbin-admin/pkg/ginx"

	"github.com/gin-gonic/gin"
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
	var req dto.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c)
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
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Update 更新角色
func (h *RoleHandler) Update(c *gin.Context) {
	var req dto.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c)
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
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Delete 删除角色
func (h *RoleHandler) Delete(c *gin.Context) {
	param := c.Param("ids")
	ids := strings.Split(param, ",")
	var idList []uint64
	for _, id := range ids {
		idInt, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			ginx.Error(c, 400, "参数错误")
			return
		}
		idList = append(idList, idInt)
	}
	if err := h.svc.Role().Delete(c, idList...); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// List 获取角色列表
func (h *RoleHandler) List(c *gin.Context) {
	var req dto.RoleListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ginx.ParamError(c)
		return
	}

	roles, total, err := h.svc.Role().List(c, req.Page, req.PageSize)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, &dto.RoleListResponse{
		List:  dto.ToRoleList(roles),
		Total: total,
	})
}

// AssignMenus 分配菜单
func (h *RoleHandler) AssignMenus(c *gin.Context) {
	var req dto.AssignMenusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c)
		return
	}

	if err := h.svc.Role().AssignMenus(c, req.RoleID, req.MenuIDs); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// GetMenus 获取角色菜单
func (h *RoleHandler) GetMenus(c *gin.Context) {
	// 从查询参数获取角色ID
	roleID := c.Query("role_id")
	if roleID == "" {
		ginx.ParamError(c)
		return
	}

	// 转换为 uint64
	id, err := strconv.ParseUint(roleID, 10, 64)
	if err != nil {
		ginx.ParamError(c)
		return
	}

	// 获取角色的菜单列表
	menus, err := h.svc.Role().GetRoleMenus(c, id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, dto.ToMenuList(menus))
}

func (h *RoleHandler) Detail(ctx *gin.Context) {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		ginx.Error(ctx, 400, "参数错误")
		return
	}
	role, err := h.svc.Role().FindByID(ctx.Request.Context(), id)
	if err != nil {
		ginx.Error(ctx, 500, err.Error())
		return
	}
	ginx.Success(ctx, dto.ToRoleResponse(role))
}
