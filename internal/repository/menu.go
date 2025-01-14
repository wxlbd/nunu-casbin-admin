package repository

import (
	"context"
	"fmt"

	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/types"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(ctx context.Context, menu *model.Menu) (uint64, error)
	Update(ctx context.Context, menu *model.Menu) error
	Delete(ctx context.Context, id ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Menu, error)
	FindByIDs(ctx context.Context, ids []uint64) ([]*model.Menu, error)
	FindByName(ctx context.Context, name ...string) ([]*model.Menu, error)
	List(ctx context.Context, query *model.MenuQuery) ([]*model.Menu, int64, error)
	FindByParentID(ctx context.Context, parentID uint64) ([]*model.Menu, error)
	FindByRoleID(ctx context.Context, roleID uint64) ([]*model.Menu, error)
	FindAll(ctx context.Context) ([]*model.Menu, error)
	BatchUpdate(ctx context.Context, menus []*model.Menu) error
	BatchCreate(ctx context.Context, menus []*model.Menu) error
}

type menuRepository struct {
	db *gorm.DB
}

func (r *menuRepository) BatchCreate(ctx context.Context, menus []*model.Menu) error {
	return r.db.WithContext(ctx).Create(&menus).Error
}

func (r *menuRepository) BatchUpdate(ctx context.Context, menus []*model.Menu) error {
	return r.db.WithContext(ctx).Save(menus).Error
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{
		db: db,
	}
}

func (r *menuRepository) Create(ctx context.Context, menu *model.Menu) (uint64, error) {
	if err := r.db.WithContext(ctx).Create(menu).Error; err != nil {
		return 0, err
	}
	return menu.ID, nil
}

func (r *menuRepository) Update(ctx context.Context, menu *model.Menu) error {
	return r.db.WithContext(ctx).Updates(menu).Error
}

func (r *menuRepository) Delete(ctx context.Context, id ...uint64) error {
	return r.db.WithContext(ctx).Where("id IN ?", id).Delete(&model.Menu{}).Error
}

func (r *menuRepository) FindByID(ctx context.Context, id uint64) (*model.Menu, error) {
	var menu model.Menu
	err := r.db.WithContext(ctx).First(&menu, id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) List(ctx context.Context, query *model.MenuQuery) ([]*model.Menu, int64, error) {
	var menus []*model.Menu
	var total int64

	db := r.db.WithContext(ctx)

	// 如果查询参数为空，设置默认值
	if query == nil {
		query = &model.MenuQuery{
			PageParam: types.PageParam{
				Page:     1,
				PageSize: 10,
			},
		}
	}
	// 构建查询条件
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Path != "" {
		db = db.Where("path LIKE ?", "%"+query.Path+"%")
	}
	if query.Component != "" {
		db = db.Where("component LIKE ?", "%"+query.Component+"%")
	}
	if query.Status != 0 {
		db = db.Where("status = ?", query.Status)
	}

	// 排序
	if query.OrderBy != "" {
		order := "ASC"
		db = db.Order(fmt.Sprintf("%s %s", query.OrderBy, order))
	} else {
		db = db.Order("sort ASC, id DESC")
	}

	// 统计总数
	err := db.Model(&model.Menu{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (query.Page - 1) * query.PageSize
	err = db.Offset(offset).Limit(query.PageSize).Find(&menus).Error
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

func (r *menuRepository) FindByName(ctx context.Context, names ...string) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := r.db.WithContext(ctx).Model(&model.Menu{}).Where("name IN ?", names).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) FindByIDs(ctx context.Context, ids []uint64) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&menus).Error
	return menus, err
}

func (r *menuRepository) FindAll(ctx context.Context) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := r.db.WithContext(ctx).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}
