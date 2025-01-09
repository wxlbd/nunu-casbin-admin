package service

import (
	"context"
	"errors"

	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/repository"
)

type MenuService interface {
	Create(ctx context.Context, menu *model.Menu) error
	Update(ctx context.Context, menu *model.Menu) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Menu, error)
	List(ctx context.Context, page, size int) ([]*model.Menu, int64, error)
	GetMenuTree(ctx context.Context) ([]*MenuTree, error)
	GetUserMenus(ctx context.Context, userID uint64) ([]*MenuTree, error)
}

type MenuTree struct {
	*model.Menu
	Children []*MenuTree `json:"children"`
}

type menuService struct {
	repo repository.Repository
}

func NewMenuService(repo repository.Repository) MenuService {
	return &menuService{
		repo: repo,
	}
}

func (s *menuService) Create(ctx context.Context, menu *model.Menu) error {
	// 检查菜单名称是否已存在
	existMenu, _ := s.repo.Menu().FindByName(ctx, menu.Name)
	if existMenu != nil {
		return errors.New("菜单名称已存在")
	}

	// 如果不是根菜单，检查父菜单是否存在
	if menu.ParentID != 0 {
		parentMenu, err := s.repo.Menu().FindByID(ctx, menu.ParentID)
		if err != nil {
			return err
		}
		if parentMenu == nil {
			return errors.New("父菜单不存在")
		}
	}

	return s.repo.Menu().Create(ctx, menu)
}

func (s *menuService) Update(ctx context.Context, menu *model.Menu) error {
	existMenu, err := s.repo.Menu().FindByID(ctx, menu.ID)
	if err != nil {
		return err
	}
	if existMenu == nil {
		return errors.New("菜单不存在")
	}

	// 如果修改了菜单名称，需要检查新名称是否已存在
	if menu.Name != existMenu.Name {
		if exist, _ := s.repo.Menu().FindByName(ctx, menu.Name); exist != nil {
			return errors.New("菜单名称已存在")
		}
	}

	// 检查是否形成循环依赖
	if menu.ParentID != 0 {
		parent := menu.ParentID
		for parent != 0 {
			parentMenu, err := s.repo.Menu().FindByID(ctx, parent)
			if err != nil {
				return err
			}
			if parentMenu == nil {
				return errors.New("父菜单不存在")
			}
			if parentMenu.ID == menu.ID {
				return errors.New("不能将菜单的子菜单设为其父菜单")
			}
			parent = parentMenu.ParentID
		}
	}

	return s.repo.Menu().Update(ctx, menu)
}

func (s *menuService) Delete(ctx context.Context, id uint64) error {
	// 检查是否有子菜单
	children, err := s.repo.Menu().FindByParentID(ctx, id)
	if err != nil {
		return err
	}
	if len(children) > 0 {
		return errors.New("请先删除子菜单")
	}

	// 检查是否有角色关联此菜单
	roles, err := s.repo.RoleMenu().FindRolesByMenuID(ctx, id)
	if err != nil {
		return err
	}
	if len(roles) > 0 {
		return errors.New("该菜单已被角色使用，无法删除")
	}

	return s.repo.Menu().Delete(ctx, id)
}

func (s *menuService) FindByID(ctx context.Context, id uint64) (*model.Menu, error) {
	return s.repo.Menu().FindByID(ctx, id)
}

func (s *menuService) List(ctx context.Context, page, size int) ([]*model.Menu, int64, error) {
	return s.repo.Menu().List(ctx, page, size)
}

func (s *menuService) GetMenuTree(ctx context.Context) ([]*MenuTree, error) {
	menus, _, err := s.repo.Menu().List(ctx, 1, 1000) // 获取所有菜单
	if err != nil {
		return nil, err
	}

	return s.buildMenuTree(menus, 0), nil
}

func (s *menuService) GetUserMenus(ctx context.Context, userID uint64) ([]*MenuTree, error) {
	// 获取用户的所有角色
	roles, err := s.repo.UserRole().FindRolesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 获取所有角色的菜单
	menuMap := make(map[uint64]*model.Menu)
	for _, role := range roles {
		menus, err := s.repo.RoleMenu().FindMenusByRoleID(ctx, role.ID)
		if err != nil {
			return nil, err
		}
		for _, menu := range menus {
			menuMap[menu.ID] = menu
		}
	}

	// 转换为切片
	var menus []*model.Menu
	for _, menu := range menuMap {
		menus = append(menus, menu)
	}

	return s.buildMenuTree(menus, 0), nil
}

// 构建菜单树
func (s *menuService) buildMenuTree(menus []*model.Menu, parentID uint64) []*MenuTree {
	var trees []*MenuTree
	for _, menu := range menus {
		if menu.ParentID == parentID {
			tree := &MenuTree{
				Menu:     menu,
				Children: s.buildMenuTree(menus, menu.ID),
			}
			trees = append(trees, tree)
		}
	}
	return trees
}
