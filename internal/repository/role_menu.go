package repository

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"gorm.io/gorm"
)

type RoleMenuRepository interface {
	WithTx(tx *gorm.DB) RoleMenuRepository
	Create(ctx context.Context, roleID, menuID uint64) error
	Delete(ctx context.Context, roleID, menuID uint64) error
	DeleteByRoleID(ctx context.Context, roleID uint64) error
	FindMenusByRoleID(ctx context.Context, roleID uint64) ([]*model.Menu, error)
	FindRolesByMenuID(ctx context.Context, menuID uint64) ([]*model.Role, error)
	BatchCreate(ctx context.Context, roleID uint64, menuIDs []uint64) error
	FindMenusByRoleIDs(ctx context.Context, roleIDs ...uint64) ([]*model.Menu, error)
	FindRolesByMenuIDs(ctx context.Context, menuIDs []uint64) ([]*model.Role, error)
}

type roleMenuRepository struct {
	db *gorm.DB
}

func NewRoleMenuRepository(db *gorm.DB) RoleMenuRepository {
	return &roleMenuRepository{
		db: db,
	}
}

func (r *roleMenuRepository) WithTx(tx *gorm.DB) RoleMenuRepository {
	return &roleMenuRepository{
		db: tx,
	}
}

func (r *roleMenuRepository) Create(ctx context.Context, roleID, menuID uint64) error {
	return r.db.WithContext(ctx).Create(map[string]interface{}{
		"role_id": roleID,
		"menu_id": menuID,
	}).Error
}

func (r *roleMenuRepository) Delete(ctx context.Context, roleID, menuID uint64) error {
	return r.db.WithContext(ctx).
		Where("role_id = ? AND menu_id = ?", roleID, menuID).
		Delete(&model.RoleMenus{}).Error
}

func (r *roleMenuRepository) DeleteByRoleID(ctx context.Context, roleID uint64) error {
	return r.db.WithContext(ctx).
		Where("role_id = ?", roleID).
		Delete(&model.RoleMenus{}).Error
}

func (r *roleMenuRepository) FindMenusByRoleID(ctx context.Context, roleID uint64) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := r.db.WithContext(ctx).
		Joins("JOIN role_menus ON role_menus.menu_id = menu.id").
		Where("role_menus.role_id = ?", roleID).
		Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *roleMenuRepository) FindMenusByRoleIDs(ctx context.Context, roleIDs ...uint64) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := r.db.WithContext(ctx).
		Joins("JOIN role_menus ON role_menus.menu_id = menu.id").
		Where("role_menus.role_id IN ?", roleIDs).
		Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *roleMenuRepository) FindRolesByMenuID(ctx context.Context, menuID uint64) ([]*model.Role, error) {
	var roles []*model.Role
	err := r.db.WithContext(ctx).
		Joins("JOIN role_menus ON role_menus.role_id = role.id").
		Where("role_menus.menu_id = ?", menuID).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleMenuRepository) BatchCreate(ctx context.Context, roleID uint64, menuIDs []uint64) error {
	// 开启事务
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先删除原有的关联
		if err := tx.Where("role_id = ?", roleID).Delete(&model.RoleMenus{}).Error; err != nil {
			return err
		}

		// 批量创建新的关联
		var roleMenus []map[string]interface{}
		for _, menuID := range menuIDs {
			roleMenus = append(roleMenus, map[string]interface{}{
				"role_id": roleID,
				"menu_id": menuID,
			})
		}
		return tx.Model(&model.RoleMenus{}).Create(roleMenus).Error
	})
}

func (r *roleMenuRepository) FindRolesByMenuIDs(ctx context.Context, menuIDs []uint64) ([]*model.Role, error) {
	var roles []*model.Role
	err := r.db.WithContext(ctx).
		Joins("JOIN role_menus ON role_menus.role_id = role.id").
		Where("role_menus.menu_id IN ?", menuIDs).
		Find(&roles).Error
	return roles, err
}
