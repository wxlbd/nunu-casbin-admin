package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wxlbd/nunu-casbin-admin/internal/dto"
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/service"
	"github.com/wxlbd/nunu-casbin-admin/pkg/ginx"
)

type MenuHandler struct {
	svc service.Service
}

func NewMenuHandler(svc service.Service) *MenuHandler {
	return &MenuHandler{
		svc: svc,
	}
}

// Create 创建菜单
func (h *MenuHandler) Create(c *gin.Context) {
	var req dto.MenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c)
		return
	}

	menu := &model.Menu{
		ParentID:  req.ParentID,
		Name:      req.Name,
		Path:      req.Path,
		Component: req.Component,
		Sort:      req.Sort,
		Status:    req.Status,
	}

	if err := h.svc.Menu().Create(c, menu); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Update 更新菜单
func (h *MenuHandler) Update(c *gin.Context) {
	var req dto.MenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c)
		return
	}

	menu := &model.Menu{
		ID:        req.ID,
		ParentID:  req.ParentID,
		Name:      req.Name,
		Path:      req.Path,
		Component: req.Component,
		Sort:      req.Sort,
		Status:    req.Status,
	}

	if err := h.svc.Menu().Update(c, menu); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Delete 删除菜单
func (h *MenuHandler) Delete(c *gin.Context) {
	var req dto.MenuIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c)
		return
	}

	if err := h.svc.Menu().Delete(c, req.ID); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// GetMenuTree 获取菜单树
func (h *MenuHandler) GetMenuTree(c *gin.Context) {
	tree, err := h.svc.Menu().GetMenuTree(c)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, dto.ToMenuTree(tree))
}

// GetUserMenus 获取用户菜单
func (h *MenuHandler) GetUserMenus(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		ginx.Unauthorized(c)
		return
	}

	// 获取用户的角色
	roles, err := h.svc.User().GetUserRoles(c, userID)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	// 检查是否有超级管理员角色
	isAdmin := false
	for _, role := range roles {
		if role.Code == "SuperAdmin" {
			isAdmin = true
			break
		}
	}

	var menus []*service.MenuTree
	if isAdmin {
		menus, err = h.svc.Menu().GetMenuTree(c)
	} else {
		menus, err = h.svc.Menu().GetUserMenus(c, userID)
	}

	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, dto.ToMenuTree(menus))
}
