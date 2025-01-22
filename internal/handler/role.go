package handler

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wxlbd/gin-casbin-admin/internal/dto"
	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"github.com/wxlbd/gin-casbin-admin/pkg/ginx"
)

type RoleHandler struct {
	svc Service
}

func NewRoleHandler(svc Service) *RoleHandler {
	return &RoleHandler{
		svc: svc,
	}
}

// Create 创建角色
// @Summary 创建角色
// @Description 创建一个新的角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param data body dto.RoleRequest true "角色信息"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/role [post]
func (h *RoleHandler) Create(c *gin.Context) {
	var req dto.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
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
// @Summary 更新角色
// @Description 更新指定ID的角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param data body dto.RoleRequest true "角色信息"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "角色不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/role/{id} [put]
func (h *RoleHandler) Update(c *gin.Context) {
	var req dto.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
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
// @Summary 删除角色
// @Description 删除指定ID的角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param ids path string true "角色ID列表(多个用逗号分隔)"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/role/{ids} [delete]
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
// @Summary 获取角色列表
// @Description 分页获取角色列表
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param name query string false "角色名称"
// @Param code query string false "角色编码"
// @Param status query int false "状态(1:正常 2:禁用)"
// @Success 200 {object} ginx.Response{data=ginx.ListData{list=[]dto.RoleResponse,total=int64}} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/role [get]
func (h *RoleHandler) List(c *gin.Context) {
	var req dto.RoleListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	roles, total, err := h.svc.Role().List(c, &req)
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
// @Summary 分配菜单
// @Description 为指定角色分配菜单权限
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param data body dto.AssignMenusRequest true "菜单权限列表"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "角色不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/role/{id}/menus [post]
func (h *RoleHandler) AssignMenus(c *gin.Context) {
	var req dto.AssignMenusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	param := c.Param("id")
	req.RoleID, _ = strconv.ParseUint(param, 10, 64)
	if err := h.svc.Role().AssignMenus(c, req.RoleID, req.Permissions); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// GetPermittedMenus 获取角色菜单
// @Summary 获取角色菜单
// @Description 获取指定角色的菜单权限列表
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} ginx.Response{data=[]dto.Menu} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "角色不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/role/{id}/menus [get]
func (h *RoleHandler) GetPermittedMenus(c *gin.Context) {
	// 从查询参数获取角色ID
	roleID := c.Param("id")
	if roleID == "" {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "角色ID不能为空"))
		return
	}

	// 转换为 uint64
	id, err := strconv.ParseUint(roleID, 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的角色ID"))
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

// Detail 获取角色详情
// @Summary 获取角色详情
// @Description 获取指定ID的角色详情
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} ginx.Response{data=dto.RoleResponse} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "角色不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/role/{id} [get]
func (h *RoleHandler) Detail(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的角色ID"))
		return
	}
	role, err := h.svc.Role().FindByID(c.Request.Context(), id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, dto.ToRoleResponse(role))
}
