package service

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/handler"
	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
)

type sysMenuService struct {
	repo Repository
}

func NewSysMenuService(repo Repository) handler.SysMenuService {
	return &sysMenuService{
		repo: repo,
	}
}

func (s *sysMenuService) Create(ctx context.Context, menu *model.SysMenu) error {
	// 1. 检查父菜单是否存在
	if menu.ParentID != 0 {
		parent, err := s.repo.SysMenu().Get(ctx, menu.ParentID)
		if err != nil {
			return err
		}
		if parent == nil {
			return errors.WithMsg(errors.NotFound, "父菜单不存在")
		}
	}

	// 2. 检查菜单名称是否已存在
	exist, err := s.repo.SysMenu().FindByTitle(ctx, menu.Title)
	if err != nil {
		return err
	}
	if exist != nil {
		return errors.WithMsg(errors.AlreadyExists, "菜单名称已存在")
	}

	// 3. 创建菜单
	return s.repo.SysMenu().Create(ctx, menu)
}

func (s *sysMenuService) Update(ctx context.Context, menu *model.SysMenu) error {
	// 1. 检查菜单是否存在
	old, err := s.repo.SysMenu().Get(ctx, menu.ID)
	if err != nil {
		return err
	}
	if old == nil {
		return errors.WithMsg(errors.NotFound, "菜单不存在")
	}

	// 2. 检查父菜单是否存在且不会形成循环依赖
	if menu.ParentID != 0 {
		parent := menu.ParentID
		for parent != 0 {
			parentMenu, err := s.repo.SysMenu().Get(ctx, parent)
			if err != nil {
				return err
			}
			if parentMenu == nil {
				return errors.WithMsg(errors.NotFound, "父菜单不存在")
			}
			if parentMenu.ID == menu.ID {
				return errors.WithMsg(errors.InvalidParam, "菜单形成循环依赖")
			}
			parent = parentMenu.ParentID
		}
	}

	// 3. 如果修改了菜单名称，检查新名称是否已存在
	if menu.Title != old.Title {
		exist, err := s.repo.SysMenu().FindByTitle(ctx, menu.Title)
		if err != nil {
			return err
		}
		if exist != nil && exist.ID != menu.ID {
			return errors.WithMsg(errors.AlreadyExists, "菜单名称已存在")
		}
	}

	// 4. 更新菜单
	return s.repo.SysMenu().Save(ctx, menu)
}

func (s *sysMenuService) Delete(ctx context.Context, ids ...int64) error {
	// 1. 检查是否有子菜单
	for _, id := range ids {
		children, err := s.repo.SysMenu().FindByParentID(ctx, id)
		if err != nil {
			return err
		}
		if len(children) > 0 {
			return errors.WithMsg(errors.InvalidParam, "存在子菜单，无法删除")
		}
	}

	// 2. 删除菜单
	return s.repo.SysMenu().Delete(ctx, ids...)
}

func (s *sysMenuService) Get(ctx context.Context, id int64) (*model.SysMenu, error) {
	return s.repo.SysMenu().Get(ctx, id)
}

func (s *sysMenuService) List(ctx context.Context, query *model.SysMenuQuery) ([]*model.SysMenu, int64, error) {
	return s.repo.SysMenu().List(ctx, query)
}

// 获取所有菜单
func (s *sysMenuService) GetAllMenus(ctx context.Context) ([]*model.SysMenu, error) {
	return s.repo.SysMenu().FindAll(ctx)
}

// buildTree 构建菜单树
func buildTree(menus []*model.SysMenu, parentID int64) []*model.SysMenuTree {
	var trees []*model.SysMenuTree
	for _, menu := range menus {
		if menu.ParentID == parentID {
			tree := &model.SysMenuTree{
				SysMenu:  menu,
				Children: buildTree(menus, menu.ID),
			}
			trees = append(trees, tree)
		}
	}
	return trees
}

func (s *sysMenuService) GetMenuTree(ctx context.Context) ([]*model.SysMenuTree, error) {
	menus, err := s.repo.SysMenu().FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return buildTree(menus, 0), nil
}

func (s *sysMenuService) GetUserMenuTree(ctx context.Context, userID uint64) ([]*model.SysMenuTree, error) {
	// 1. 获取用户角色
	roles, err := s.repo.UserRole().FindRolesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 2. 如果是超级管理员,返回所有菜单
	for _, role := range roles {
		if role.Code == "SuperAdmin" {
			menus, err := s.repo.SysMenu().FindAll(ctx)
			if err != nil {
				return nil, err
			}
			return buildTree(menus, 0), nil
		}
	}

	// 3. 获取角色ID
	var roleIDs []uint64
	for _, role := range roles {
		roleIDs = append(roleIDs, role.ID)
	}

	// 4. 获取角色菜单
	menus, err := s.repo.SysMenu().FindByRoleIDs(ctx, roleIDs...)
	if err != nil {
		return nil, err
	}

	return buildTree(menus, 0), nil
}
