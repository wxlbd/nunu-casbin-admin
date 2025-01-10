package dto

import (
	"time"

	"github.com/wxlbd/nunu-casbin-admin/internal/handler/response"
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/service"
)

func ToMenuList(menus []*model.Menu) []*response.Menu {
	var menuList []*response.Menu
	for _, menu := range menus {
		menuList = append(menuList, ToMenu(menu))
	}
	return menuList
}

func ToMenu(menu *model.Menu) *response.Menu {
	if menu == nil {
		return nil
	}
	return &response.Menu{
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

func ToMenuTree(menus []*service.MenuTree) []*response.MenuTreeNode {
	var menuList []*response.MenuTreeNode
	for _, menu := range menus {
		menuList = append(menuList, ToMenuTreeNode(menu))
	}
	return menuList
}

func ToMenuTreeNode(menu *service.MenuTree) *response.MenuTreeNode {
	if menu == nil {
		return nil
	}
	return &response.MenuTreeNode{
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
func ToMenuMeta(meta model.MenuMeta) response.Meta {
	return response.Meta{
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
