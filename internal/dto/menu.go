package dto

import (
	"time"

	"github.com/wxlbd/nunu-casbin-admin/internal/types"

	"github.com/wxlbd/nunu-casbin-admin/internal/model"
)

func ToMenuList(menus []*model.Menu) []*Menu {
	var menuList []*Menu
	for _, menu := range menus {
		menuList = append(menuList, ToMenu(menu))
	}
	return menuList
}

func ToMenu(menu *model.Menu) *Menu {
	if menu == nil {
		return nil
	}
	return &Menu{
		ID:         menu.ID,
		ParentID:   menu.ParentID,
		Name:       menu.Name,
		Path:       menu.Path,
		Component:  menu.Component,
		Sort:       menu.Sort,
		Status:     menu.Status,
		CreateTime: menu.CreatedAt.Format(time.DateTime),
		UpdateTime: menu.UpdatedAt.Format(time.DateTime),
	}
}

func ToMenuTree(menus []*model.MenuTree) []*MenuTreeNode {
	var menuList []*MenuTreeNode
	for _, menu := range menus {
		menuList = append(menuList, ToMenuTreeNode(menu))
	}
	return menuList
}

func ToMenuTreeNode(menu *model.MenuTree) *MenuTreeNode {
	if menu == nil {
		return nil
	}
	return &MenuTreeNode{
		ID:        menu.ID,
		ParentID:  menu.ParentID,
		Name:      menu.Name,
		Path:      menu.Path,
		Component: menu.Component,
		Meta:      menu.Meta,
		Sort:      menu.Sort,
		Status:    menu.Status,
		CreatedAt: menu.CreatedAt.Format(time.DateTime),
		UpdatedAt: menu.UpdatedAt.Format(time.DateTime),
		Children:  ToMenuTree(menu.Children),
	}
}

// MenuTreeNode MenuTree 菜单树节点
type MenuTreeNode struct {
	ID        uint64          `json:"id"`
	ParentID  uint64          `json:"parent_id"`
	Name      string          `json:"name"`
	Meta      *types.MenuMeta `json:"meta"`
	Path      string          `json:"path"`
	Component string          `json:"component"`
	Redirect  string          `json:"redirect"`
	Status    int8            `json:"status"`
	Sort      int16           `json:"sort"`
	CreatedBy int             `json:"created_by"`
	UpdatedBy int             `json:"updated_by"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	Remark    string          `json:"remark"`
	Children  []*MenuTreeNode `json:"children"`
}

// MenuTreeResponse 菜单列表响应
type MenuTreeResponse []*MenuTreeNode

// Menu 菜单信息响应
type Menu struct {
	ID         uint64 `json:"id"`
	ParentID   uint64 `json:"parent_id"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	Sort       int16  `json:"sort"`
	Status     int8   `json:"status"`
	Mate       any    `json:"mate"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// MenuRequest 创建/更新菜单请求
type MenuRequest struct {
	ID        uint64 `json:"id"` // 更新时必填
	ParentID  uint64 `json:"parent_id"`
	Name      string `json:"name" binding:"required"`
	Path      string `json:"path" binding:"required"`
	Component string `json:"component" binding:"required"`
	Sort      int16  `json:"sort"`
	Status    int8   `json:"status"`
}

// MenuIDRequest 删除菜单请求
type MenuIDRequest struct {
	ID uint64 `json:"id" binding:"required"`
}

// MenuListRequest 菜单列表请求
type MenuListRequest struct {
	types.PageParam
	Name      string `form:"name"`
	Path      string `form:"path"`
	Component string `form:"component"`
	Status    int8   `form:"status"`
}

func (req *MenuListRequest) ToModel() *model.MenuQuery {
	return &model.MenuQuery{
		Name:      req.Name,
		Path:      req.Path,
		Component: req.Component,
		Status:    req.Status,
		PageParam: types.PageParam{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}
}

type BtnPermission struct {
	ID    uint64 `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
	I18N  string `json:"i18n"`
	Type  string `json:"type"`
}

type MenuBase struct {
	ID             uint64          `json:"id"`
	ParentId       uint64          `json:"parent_id"`
	Name           string          `json:"name"`
	Path           string          `json:"path"`
	Meta           *types.MenuMeta `json:"meta"`
	Component      string          `json:"component"`
	Sort           int16           `json:"sort"`
	Status         int8            `json:"status"`
	BtnPermissions []BtnPermission `json:"btnPermission"`
	Redirect       string          `json:"redirect"`
	Remark         string          `json:"remark"`
	Title          string          `json:"title"`
}

type CreateMenuRequest struct {
	MenuBase
	CreatedBy uint64 `json:"created_by"`
}

func (c *CreateMenuRequest) ToModel() *model.Menu {
	return &model.Menu{
		ParentID:  c.ParentId,
		Name:      c.Name,
		Path:      c.Path,
		Meta:      c.Meta,
		Component: c.Component,
		Sort:      c.Sort,
		Status:    c.Status,
		CreatedBy: c.CreatedBy,
		Remark:    c.Remark,
		Redirect:  c.Redirect,
	}
}

func (c *MenuBase) BtnPermissionsToModels() []*model.Menu {
	var models []*model.Menu
	for _, btn := range c.BtnPermissions {
		models = append(models, &model.Menu{
			ID:       btn.ID,
			Name:     btn.Code,
			ParentID: c.ID,
			Sort:     0,
			Status:   1,
			Meta: &types.MenuMeta{
				Title: btn.Title,
				I18n:  btn.I18N,
				Type:  "B",
			},
		})
	}
	return models
}

type UpdateMenuRequest struct {
	MenuBase
	UpdatedBy int `json:"updated_by"`
}

func (c *UpdateMenuRequest) ToModel() *model.Menu {
	return &model.Menu{
		ID:        c.ID,
		ParentID:  c.ParentId,
		Name:      c.Name,
		Path:      c.Path,
		Meta:      c.Meta,
		Component: c.Component,
		Sort:      c.Sort,
		Status:    c.Status,
		UpdatedBy: uint64(c.UpdatedBy),
		Remark:    c.Remark,
		Redirect:  c.Redirect,
	}
}
