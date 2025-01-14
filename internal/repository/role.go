package repository

import (
	"context"
	"errors"

	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(ctx context.Context, role *model.Role) error
	Update(ctx context.Context, role *model.Role) error
	Delete(ctx context.Context, id ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Role, error)
	FindByCode(ctx context.Context, code string) (*model.Role, error)
	List(ctx context.Context, page, size int) ([]*model.Role, int64, error)
	// FindByIDs 根据角色ID列表查询角色
	FindByIDs(ctx context.Context, ids []uint64) ([]*model.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

// FindByIDs implements RoleRepository.
func (r *roleRepository) FindByIDs(ctx context.Context, ids []uint64) ([]*model.Role, error) {
	var roles []*model.Role
	err := r.db.WithContext(ctx).Where("id IN (?)", ids).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) Create(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *roleRepository) Update(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Updates(role).Error
}

func (r *roleRepository) Delete(ctx context.Context, id ...uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Role{}, id).Error
}

func (r *roleRepository) FindByID(ctx context.Context, id uint64) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) FindByCode(ctx context.Context, code string) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) List(ctx context.Context, page, size int) ([]*model.Role, int64, error) {
	var roles []*model.Role
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Role{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err = r.db.WithContext(ctx).Offset(offset).Limit(size).Find(&roles).Error
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}
