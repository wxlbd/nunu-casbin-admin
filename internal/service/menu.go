package service

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/pkg/errors"

	"github.com/wxlbd/gin-casbin-admin/internal/dto"

	"github.com/casbin/casbin/v2"
	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/internal/types"
)

type MenuService interface {
	Create(ctx context.Context, req *dto.CreateMenuRequest) error
	Update(ctx context.Context, req *dto.UpdateMenuRequest) error
	Delete(ctx context.Context, id ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Menu, error)
	List(ctx context.Context, query *model.MenuQuery) ([]*model.Menu, int64, error)
	GetMenuTree(ctx context.Context) ([]*model.MenuTree, error)
	GetUserMenus(ctx context.Context, userID uint64) ([]*model.MenuTree, error)
	GetAllMenus(ctx context.Context) ([]*model.Menu, error)
}

type menuService struct {
	repo     Repository
	enforcer *casbin.Enforcer
}

func NewMenuService(repo Repository, enforcer *casbin.Enforcer) MenuService {
	return &menuService{
		repo:     repo,
		enforcer: enforcer,
	}
}

func (s *menuService) Create(ctx context.Context, req *dto.CreateMenuRequest) error {
	// 1. 转换请求为菜单模型
	menu := req.ToModel()

	// 2. 查询菜单名称是否已存在
	if ms, _ := s.repo.Menu().FindByNames(ctx, menu.Name); len(ms) > 0 {
		return errors.WithMsg(errors.AlreadyExists, "菜单名称已存在")
	}

	err := s.repo.Menu().Create(ctx, menu)
	if err != nil {
		return err
	}
	var menus []*model.Menu
	if len(req.BtnPermissions) > 0 {
		menus = append(menus, req.BtnPermissionsToModels()...)
	}
	for i := range menus {
		menus[i].ParentID = menu.ID
	}
	return s.repo.Menu().BatchCreate(ctx, menus)
}

func (s *menuService) Update(ctx context.Context, req *dto.UpdateMenuRequest) error {
	// 1. 转换请求为菜单模型
	menus := []*model.Menu{req.ToModel()}
	if len(req.BtnPermissions) > 0 {
		menus = append(menus, req.BtnPermissionsToModels()...)
	}

	// 2. 查询旧的菜单信息
	oldMenu, err := s.repo.Menu().FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if oldMenu == nil {
		return errors.WithMsg(errors.NotFound, "菜单不存在")
	}

	// 3. 如果修改了菜单名称，需要检查新名称是否已存在
	if req.Name != oldMenu.Name {
		if ms, _ := s.repo.Menu().FindByNames(ctx, req.Name); len(ms) > 0 {
			return errors.WithMsg(errors.AlreadyExists, "菜单名称已存在")
		}
	}

	// 4. 检查是否形成循环依赖
	if req.ParentId != 0 {
		parent := req.ParentId
		for parent != 0 {
			parentMenu, err := s.repo.Menu().FindByID(ctx, parent)
			if err != nil {
				return err
			}
			if parentMenu == nil {
				return errors.WithMsg(errors.NotFound, "父菜单不存在")
			}
			if parentMenu.ID == req.ID {
				return errors.WithMsg(errors.InvalidParam, "菜单形成循环依赖")
			}
			parent = parentMenu.ParentID
		}
	}

	// 5. 更新权限策略
	// 5.1 获取所有需要更新的菜单ID（包括子菜单）
	menuIDs := []uint64{req.ID}
	for _, menu := range menus {
		if menu.ID != 0 {
			menuIDs = append(menuIDs, menu.ID)
		}
	}

	// 5.2 查询这些菜单的旧数据
	oldMenus, err := s.repo.Menu().FindByIDs(ctx, menuIDs)
	if err != nil {
		return err
	}

	// 5.3 获取拥有这些菜单的所有角色
	roles, err := s.repo.RoleMenu().FindRolesByMenuIDs(ctx, menuIDs)
	if err != nil {
		return err
	}

	// 5.4 更新权限策略
	for _, role := range roles {
		// 删除旧的权限策略
		for _, oldMenu := range oldMenus {
			if oldMenu.Meta != nil && oldMenu.Meta.Type == "B" {
				oldPath, oldMethod := convertMenuToAPI(oldMenu.Name)
				_, err = s.enforcer.RemovePolicy(role.Code, oldPath, oldMethod)
				if err != nil {
					return err
				}
			}
		}

		// 添加新的权限策略
		for _, menu := range menus {
			if menu.Meta != nil && menu.Meta.Type == "B" {
				newPath, newMethod := convertMenuToAPI(menu.Name)
				_, err = s.enforcer.AddPolicy(role.Code, newPath, newMethod)
				if err != nil {
					return err
				}
			}
		}
	}

	// 6. 批量更新菜单
	return s.repo.Menu().BatchUpdate(ctx, menus)
}

func (s *menuService) Delete(ctx context.Context, ids ...uint64) error {
	menus, err := s.repo.Menu().FindByIDs(ctx, ids)
	if err != nil {
		return err
	}
	for _, menu := range menus {
		// 如果是按钮类型，需要删除相关的权限策略
		if menu.Meta.Type == "B" {
			roles, err := s.repo.RoleMenu().FindRolesByMenuID(ctx, menu.ID)
			if err != nil {
				return err
			}

			path, method := convertMenuToAPI(menu.Name)
			for _, role := range roles {
				_, err = s.enforcer.RemovePolicy(role.Code, path, method)
				if err != nil {
					return err
				}
			}
		}
	}

	return s.repo.Menu().Delete(ctx, ids...)
}

func (s *menuService) FindByID(ctx context.Context, id uint64) (*model.Menu, error) {
	return s.repo.Menu().FindByID(ctx, id)
}

func (s *menuService) List(ctx context.Context, query *model.MenuQuery) ([]*model.Menu, int64, error) {
	return s.repo.Menu().List(ctx, query)
}

func (s *menuService) GetMenuTree(ctx context.Context) ([]*model.MenuTree, error) {
	menus, _, err := s.repo.Menu().List(ctx, &model.MenuQuery{
		PageParam: types.PageParam{
			Page:     1,
			PageSize: 1000,
		},
	}) // 获取所有菜单
	if err != nil {
		return nil, err
	}

	return s.buildMenuTree(menus, 0), nil
}

func (s *menuService) GetUserMenus(ctx context.Context, userID uint64) ([]*model.MenuTree, error) {
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

func (s *menuService) GetAllMenus(ctx context.Context) ([]*model.Menu, error) {
	// 直接从数据库获取所有菜单
	return s.repo.Menu().FindAll(ctx)
}

// 构建菜单树
func (s *menuService) buildMenuTree(menus []*model.Menu, parentID uint64) []*model.MenuTree {
	var trees []*model.MenuTree
	for _, menu := range menus {
		if menu.ParentID == parentID {
			tree := &model.MenuTree{
				Menu:     menu,
				Children: s.buildMenuTree(menus, menu.ID),
			}
			trees = append(trees, tree)
		}
	}
	return trees
}
