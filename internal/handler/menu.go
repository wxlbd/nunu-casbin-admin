package handler

import (
	"strconv"
	"strings"

	"github.com/wxlbd/nunu-casbin-admin/pkg/errors"

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
	var req dto.CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	if err := h.svc.Menu().Create(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Update 更新菜单
func (h *MenuHandler) Update(c *gin.Context) {
	var req dto.UpdateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	if err := h.svc.Menu().Update(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Delete 删除菜单
func (h *MenuHandler) Delete(c *gin.Context) {
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
	if err := h.svc.Menu().Delete(c, idList...); err != nil {
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
		ginx.ServerError(c, errors.WithMsg(errors.Unauthorized, "用户未登录"))
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

	var menus []*model.MenuTree
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

// // List 获取菜单列表
// func (h *MenuHandler) List(c *gin.Context) {
// 	var req dto.MenuListRequest
// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		ginx.Error(c, 400, "参数错误")
// 		return
// 	}

// 	// 参数验证
// 	if req.Page < 1 {
// 		req.Page = 1
// 	}
// 	if req.PageSize < 1 || req.PageSize > 100 {
// 		req.PageSize = 10
// 	}

// 	menus, total, err := h.svc.Menu().List(c, req.ToModel())
// 	if err != nil {
// 		ginx.Error(c, 500, err.Error())
// 		return
// 	}

// 	ginx.Success(c, &dto.MenuListResponse{
// 		List:  dto.ToMenuList(menus),
// 		Total: total,
// 	})
// }
