package repository

import (
	"context"

	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(ctx context.Context, menu *model.Menu) error
	Update(ctx context.Context, menu *model.Menu) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Menu, error)
	FindByIDs(ctx context.Context, ids []uint64) ([]*model.Menu, error)
	FindByName(ctx context.Context, name string) (*model.Menu, error)
	List(ctx context.Context, page, size int) ([]*model.Menu, int64, error)
	FindByParentID(ctx context.Context, parentID uint64) ([]*model.Menu, error)
	FindByRoleID(ctx context.Context, roleID uint64) ([]*model.Menu, error)
	FindAll(ctx context.Context) ([]*model.Menu, error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{
		db: db,
	}
}

func (r *menuRepository) Create(ctx context.Context, menu *model.Menu) error {
	return r.db.WithContext(ctx).Create(menu).Error
}

func (r *menuRepository) Update(ctx context.Context, menu *model.Menu) error {
	return r.db.WithContext(ctx).Save(menu).Error
}

func (r *menuRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Menu{}, id).Error
}

func (r *menuRepository) FindByID(ctx context.Context, id uint64) (*model.Menu, error) {
	var menu model.Menu
	err := r.db.WithContext(ctx).First(&menu, id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) List(ctx context.Context, page, size int) ([]*model.Menu, int64, error) {
	var menus []*model.Menu
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Menu{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err = r.db.WithContext(ctx).Offset(offset).Limit(size).Find(&menus).Error
	if err != nil {
		return nil, 0, err
	}

	return menus, total, nil
}

func (r *menuRepository) FindByParentID(ctx context.Context, parentID uint64) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := r.db.WithContext(ctx).Where("parent_id = ?", parentID).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) FindByRoleID(ctx context.Context, roleID uint64) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := r.db.WithContext(ctx).
		Joins("JOIN role_belongs_menu ON role_belongs_menu.menu_id = menu.id").
		Where("role_belongs_menu.role_id = ?", roleID).
		Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) FindByName(ctx context.Context, name string) (*model.Menu, error) {
	var menu model.Menu
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&menu).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) FindByIDs(ctx context.Context, ids []uint64) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := r.db.WithContext(ctx).Where("id IN (?)", ids).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) FindAll(ctx context.Context) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := r.db.WithContext(ctx).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}
