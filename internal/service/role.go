package service

import (
	"context"
	"errors"

	"github.com/casbin/casbin/v2"

	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/repository"
	"gorm.io/gorm"
)

type RoleService interface {
	Create(ctx context.Context, role *model.Role) error
	Update(ctx context.Context, role *model.Role) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Role, error)
	List(ctx context.Context, page, size int) ([]*model.Role, int64, error)
	AssignMenus(ctx context.Context, roleID uint64, menuIDs []uint64) error
	GetRoleMenus(ctx context.Context, roleID uint64) ([]*model.Menu, error)
}

type roleService struct {
	repo     repository.Repository
	enforcer *casbin.Enforcer
}

func NewRoleService(repo repository.Repository) RoleService {
	return &roleService{
		repo: repo,
	}
}

func (s *roleService) Create(ctx context.Context, role *model.Role) error {
	// 检查角色代码是否已存在
	existRole, _ := s.repo.Role().FindByCode(ctx, role.Code)
	if existRole != nil {
		return errors.New("角色代码已存在")
	}

	return s.repo.Role().Create(ctx, role)
}

func (s *roleService) Update(ctx context.Context, role *model.Role) error {
	existRole, err := s.repo.Role().FindByID(ctx, role.ID)
	if err != nil {
		return err
	}
	if existRole == nil {
		return errors.New("角色不存在")
	}

	// 如果修改了角色代码，需要检查新代码是否已存在
	if role.Code != existRole.Code {
		if exist, _ := s.repo.Role().FindByCode(ctx, role.Code); exist != nil {
			return errors.New("角色代码已存在")
		}
	}

	return s.repo.Role().Update(ctx, role)
}

func (s *roleService) Delete(ctx context.Context, id uint64) error {
	role, err := s.repo.Role().FindByID(ctx, id)
	if err != nil {
		return err
	}
	if role == nil {
		return errors.New("角色不存在")
	}

	// 删除该角色的所有权限策略
	_, err = s.enforcer.DeletePermissionsForUser(role.Code)
	if err != nil {
		return err
	}

	return s.repo.Role().Delete(ctx, id)
}

func (s *roleService) FindByID(ctx context.Context, id uint64) (*model.Role, error) {
	return s.repo.Role().FindByID(ctx, id)
}

func (s *roleService) List(ctx context.Context, page, size int) ([]*model.Role, int64, error) {
	return s.repo.Role().List(ctx, page, size)
}

func (s *roleService) AssignMenus(ctx context.Context, roleID uint64, menuIDs []uint64) error {
	// 开启事务
	return s.repo.DB().Transaction(func(tx *gorm.DB) error {
		// 1. 先删除原有的角色-菜单关联
		if err := s.repo.RoleMenu().DeleteByRoleID(ctx, roleID); err != nil {
			return err
		}

		// 2. 获取所有按钮类型的菜单（即 API）
		menus, err := s.repo.Menu().FindByIDs(ctx, menuIDs)
		if err != nil {
			return err
		}

		// 3. 更新 Casbin 策略
		role, err := s.repo.Role().FindByID(ctx, roleID)
		if err != nil {
			return err
		}

		// 先删除该角色的所有策略
		_, err = s.enforcer.DeletePermissionsForUser(role.Code)
		if err != nil {
			return err
		}

		// 添加新的策略
		for _, menu := range menus {
			if menu.Meta.Type == "B" { // 按钮类型
				// 根据菜单名称生成 API 路径和方法
				// 例如：permission:user:save -> POST /api/user
				path, method := convertMenuToAPI(menu.Name)
				_, err = s.enforcer.AddPolicy(role.Code, path, method)
				if err != nil {
					return err
				}
			}
		}

		// 4. 创建新的角色-菜单关联
		return s.repo.RoleMenu().BatchCreate(ctx, roleID, menuIDs)
	})
}

// convertMenuToAPI 将菜单名称转换为 API 路径和方法
func convertMenuToAPI(menuName string) (path, method string) {
	// TODO 根据命名规范进行转换
	// 例如：
	// permission:user:save -> POST /api/permission/user
	// permission:user:update -> PUT /api/permission/user
	// permission:user:delete -> DELETE /api/permission/user
	// permission:user:index -> GET /api/permission/user
	// ...
	return path, method
}

func (s *roleService) GetRoleMenus(ctx context.Context, roleID uint64) ([]*model.Menu, error) {
	// 检查角色是否存在
	role, err := s.repo.Role().FindByID(ctx, roleID)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.New("角色不存在")
	}

	// 获取角色的菜单列表
	return s.repo.RoleMenu().FindMenusByRoleID(ctx, roleID)
}
