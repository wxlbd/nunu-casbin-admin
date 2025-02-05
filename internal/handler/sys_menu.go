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

type SysMenuHandler struct {
	svc Service
}

func NewSysMenuHandler(svc Service) *SysMenuHandler {
	return &SysMenuHandler{
		svc: svc,
	}
}

// Create 创建菜单
// @Summary 创建菜单
// @Description 创建一个新的菜单
// @Tags 系统菜单
// @Accept json
// @Produce json
// @Param data body dto.SysMenuRequest true "菜单信息"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/menu [post]
func (h *SysMenuHandler) Create(c *gin.Context) {
	var req dto.SysMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	menu := &model.SysMenu{
		ParentID:        req.ParentID,
		MenuType:        req.MenuType,
		Title:           req.Title,
		Name:            req.Name,
		Path:            req.Path,
		Component:       req.Component,
		Rank:            req.Rank,
		Redirect:        req.Redirect,
		Icon:            req.Icon,
		ExtraIcon:       req.ExtraIcon,
		EnterTransition: req.EnterTransition,
		LeaveTransition: req.LeaveTransition,
		ActivePath:      req.ActivePath,
		Auths:           req.Auths,
		FrameSrc:        req.FrameSrc,
		FrameLoading:    req.FrameLoading,
		KeepAlive:       req.KeepAlive,
		HiddenTag:       req.HiddenTag,
		FixedTag:        req.FixedTag,
		ShowLink:        req.ShowLink,
		ShowParent:      req.ShowParent,
		Status:          req.Status,
	}

	if err := h.svc.SysMenu().Create(c, menu); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Update 更新菜单
// @Summary 更新菜单
// @Description 更新指定ID的菜单
// @Tags 系统菜单
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Param data body dto.SysMenuRequest true "菜单信息"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "菜单不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/menu/{id} [put]
func (h *SysMenuHandler) Update(c *gin.Context) {
	var req dto.SysMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的菜单ID"))
		return
	}

	menu := &model.SysMenu{
		ID:              id,
		ParentID:        req.ParentID,
		MenuType:        req.MenuType,
		Title:           req.Title,
		Name:            req.Name,
		Path:            req.Path,
		Component:       req.Component,
		Rank:            req.Rank,
		Redirect:        req.Redirect,
		Icon:            req.Icon,
		ExtraIcon:       req.ExtraIcon,
		EnterTransition: req.EnterTransition,
		LeaveTransition: req.LeaveTransition,
		ActivePath:      req.ActivePath,
		Auths:           req.Auths,
		FrameSrc:        req.FrameSrc,
		FrameLoading:    req.FrameLoading,
		KeepAlive:       req.KeepAlive,
		HiddenTag:       req.HiddenTag,
		FixedTag:        req.FixedTag,
		ShowLink:        req.ShowLink,
		ShowParent:      req.ShowParent,
		Status:          req.Status,
	}

	if err := h.svc.SysMenu().Update(c, menu); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Delete 删除菜单
// @Summary 删除菜单
// @Description 删除指定ID的菜单
// @Tags 系统菜单
// @Accept json
// @Produce json
// @Param ids path string true "菜单ID列表(多个用逗号分隔)"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/menu/{ids} [delete]
func (h *SysMenuHandler) Delete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	var ids []int64
	for _, idStr := range idsStr {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的菜单ID"))
			return
		}
		ids = append(ids, id)
	}

	if err := h.svc.SysMenu().Delete(c, ids...); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// List 获取菜单列表
// @Summary 获取菜单列表
// @Description 分页获取菜单列表
// @Tags 系统菜单
// @Accept json
// @Produce json
// @Param title query string false "菜单名称"
// @Param status query int false "状态(0停用 1正常)"
// @Param type query int false "菜单类型"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(10)
// @Success 200 {object} ginx.Response{data=ginx.ListData{list=[]dto.SysMenuResponse,total=int64}} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/menu [get]
func (h *SysMenuHandler) List(c *gin.Context) {
	var req dto.SysMenuListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	list, total, err := h.svc.SysMenu().List(c, &model.SysMenuQuery{
		Title:    req.Title,
		Status:   req.Status,
		MenuType: req.MenuType,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, &ginx.ListData{
		List:  dto.ToSysMenuList(list),
		Total: total,
	})
}

// GetMenuTree 获取菜单树
// @Summary 获取菜单树
// @Description 获取所有菜单的树形结构
// @Tags 系统菜单
// @Accept json
// @Produce json
// @Success 200 {object} ginx.Response{data=[]service.MenuTree} "成功"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/menu/tree [get]
func (h *SysMenuHandler) GetMenuTree(c *gin.Context) {
	tree, err := h.svc.SysMenu().GetMenuTree(c)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, tree)
}

// GetUserMenuTree 获取用户菜单树
// @Summary 获取用户菜单树
// @Description 获取当前用户拥有权限的菜单树形结构
// @Tags 系统菜单
// @Accept json
// @Produce json
// @Success 200 {object} ginx.Response{data=[]service.MenuTree} "成功"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/menu/user-tree [get]
func (h *SysMenuHandler) GetUserMenuTree(c *gin.Context) {
	userID := c.GetUint64("user_id")
	tree, err := h.svc.SysMenu().GetUserMenuTree(c, userID)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, tree)
}
