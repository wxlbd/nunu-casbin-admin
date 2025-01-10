package dto

import (
	"github.com/wxlbd/nunu-casbin-admin/internal/handler/response"
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/service"
	"time"
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

func ToMenuTree(menus []*service.MenuTree) []*response.MenuTree {
	var menuList []*response.MenuTree
	for _, menu := range menus {
		menuList = append(menuList, ToMenuTreeItem(menu))
	}
	return menuList
}

func ToMenuTreeItem(menu *service.MenuTree) *response.MenuTree {
	if menu == nil {
		return nil
	}
	return &response.MenuTree{
		ID:         menu.ID,
		ParentID:   menu.ParentID,
		Name:       menu.Name,
		Path:       menu.Path,
		Component:  menu.Component,
		Sort:       menu.Sort,
		Status:     menu.Status,
		CreateTime: menu.CreatedAt.Format(time.DateTime),
		UpdateTime: menu.UpdatedAt.Format(time.DateTime),
		Children:   ToMenuTree(menu.Children),
	}
}
