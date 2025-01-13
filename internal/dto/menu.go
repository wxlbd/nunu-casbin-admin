package dto

import (
	"time"

	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/service"
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

func ToMenuTree(menus []*service.MenuTree) []*MenuTreeNode {
	var menuList []*MenuTreeNode
	for _, menu := range menus {
		menuList = append(menuList, ToMenuTreeNode(menu))
	}
	return menuList
}

func ToMenuTreeNode(menu *service.MenuTree) *MenuTreeNode {
	if menu == nil {
		return nil
	}
	return &MenuTreeNode{
		ID:        menu.ID,
		ParentID:  menu.ParentID,
		Name:      menu.Name,
		Path:      menu.Path,
		Component: menu.Component,
		Meta:      ToMenuMeta(menu.Meta),
		Sort:      menu.Sort,
		Status:    menu.Status,
		CreatedAt: menu.CreatedAt.Format(time.DateTime),
		UpdatedAt: menu.UpdatedAt.Format(time.DateTime),
		Children:  ToMenuTree(menu.Children),
	}
}
func ToMenuMeta(meta model.MenuMeta) Meta {
	return Meta{
		I18N:             meta.I18n,
		Icon:             meta.Icon,
		Type:             meta.Type,
		Affix:            meta.Affix,
		Cache:            meta.Cache,
		Title:            meta.Title,
		Hidden:           meta.Hidden,
		Copyright:        meta.Copyright,
		ComponentPath:    meta.ComponentPath,
		ComponentSuffix:  meta.ComponentSuffix,
		BreadcrumbEnable: meta.BreadcrumbEnable,
	}
}

// MenuTreeNode MenuTree 菜单树节点
type MenuTreeNode struct {
	ID        uint64          `json:"id"`
	ParentID  uint64          `json:"parent_id"`
	Name      string          `json:"name"`
	Meta      Meta            `json:"meta"`
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

type Meta struct {
	I18N             string `json:"i18n"`
	Icon             string `json:"icon"`
	Type             string `json:"type"`
	Affix            bool   `json:"affix"`
	Cache            bool   `json:"cache"`
	Title            string `json:"title"`
	Hidden           bool   `json:"hidden"`
	Copyright        bool   `json:"copyright"`
	ComponentPath    string `json:"componentPath"`
	ComponentSuffix  string `json:"componentSuffix"`
	BreadcrumbEnable bool   `json:"breadcrumbEnable"`
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
	Page int `form:"page" binding:"required,min=1"`
	Size int `form:"size" binding:"required,min=1,max=100"`
}
