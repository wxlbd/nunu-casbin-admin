package handler

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wxlbd/gin-casbin-admin/internal/dto"
	"github.com/wxlbd/gin-casbin-admin/pkg/ginx"
)

type MenuHandler struct {
	svc Service
}

func NewMenuHandler(svc Service) *MenuHandler {
	return &MenuHandler{
		svc: svc,
	}
}

// Create 创建菜单
// @Summary 创建菜单
// @Description 创建一个新的菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param data body dto.CreateMenuRequest true "菜单信息"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/menu [post]
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
// @Summary 更新菜单
// @Description 更新指定ID的菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param data body dto.UpdateMenuRequest true "菜单信息"
// @Param id path int true "菜单ID"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "菜单不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/menu/{id} [put]
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
// @Summary 删除菜单
// @Description 删除指定ID的菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param ids path string true "菜单ID列表(多个用逗号分隔)"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/menu/{ids} [delete]
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
// @Summary 获取菜单树
// @Description 获取所有菜单的树形结构
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Success 200 {object} ginx.Response{data=[]dto.MenuTreeResponse} "成功"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /permission/menu/tree [get]
func (h *MenuHandler) GetMenuTree(c *gin.Context) {
	tree, err := h.svc.Menu().GetMenuTree(c)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, dto.ToMenuTree(tree))
}
