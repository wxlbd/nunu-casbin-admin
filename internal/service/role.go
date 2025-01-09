package service

import (
	"context"
	"errors"
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/repository"
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
	repo repository.Repository
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
	// 检查是否有用户关联此角色
	users, err := s.repo.UserRole().FindUsersByRoleID(ctx, id)
	if err != nil {
		return err
	}
	if len(users) > 0 {
		return errors.New("该角色下存在用户，无法删除")
	}

	// 删除角色的同时需要删除角色-菜单关联
	if err := s.repo.RoleMenu().DeleteByRoleID(ctx, id); err != nil {
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
	// 检查角色是否存在
	role, err := s.repo.Role().FindByID(ctx, roleID)
	if err != nil {
		return err
	}
	if role == nil {
		return errors.New("角色不存在")
	}

	// 检查菜单是否都存在
	for _, menuID := range menuIDs {
		menu, err := s.repo.Menu().FindByID(ctx, menuID)
		if err != nil {
			return err
		}
		if menu == nil {
			return errors.New("菜单不存在")
		}
	}

	// 批量分配菜单
	return s.repo.RoleMenu().BatchCreate(ctx, roleID, menuIDs)
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

	return s.repo.RoleMenu().FindMenusByRoleID(ctx, roleID)
}
