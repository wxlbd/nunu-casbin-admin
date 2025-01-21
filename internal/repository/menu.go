package repository

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/internal/types"
)

type MenuRepository interface {
	WithTx(tx *Query) MenuRepository
	Create(ctx context.Context, menu *model.Menu) error
	BatchCreate(ctx context.Context, menus []*model.Menu) error
	Update(ctx context.Context, menu *model.Menu) error
	BatchUpdate(ctx context.Context, menus []*model.Menu) error
	Delete(ctx context.Context, ids ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Menu, error)
	FindByIDs(ctx context.Context, ids []uint64) ([]*model.Menu, error)
	FindByNames(ctx context.Context, names ...string) ([]*model.Menu, error)
	FindAll(ctx context.Context) ([]*model.Menu, error)
	List(ctx context.Context, query *model.MenuQuery) ([]*model.Menu, int64, error)
	FindByParentID(ctx context.Context, parentID uint64) ([]*model.Menu, error)
	FindByRoleID(ctx context.Context, roleID uint64) ([]*model.Menu, error)
}

type menuRepository struct {
	query *Query
}

func NewMenuRepository(query *Query) MenuRepository {
	return &menuRepository{query: query}
}

func (r *menuRepository) WithTx(tx *Query) MenuRepository {
	return &menuRepository{query: tx}
}

func (r *menuRepository) Create(ctx context.Context, menu *model.Menu) error {
	return r.query.WithContext(ctx).Menu.Create(menu)
}

func (r *menuRepository) BatchCreate(ctx context.Context, menus []*model.Menu) error {
	return r.query.WithContext(ctx).Menu.Create(menus...)
}

func (r *menuRepository) Update(ctx context.Context, menu *model.Menu) error {
	_, err := r.query.Menu.WithContext(ctx).Updates(menu)
	if err != nil {
		return err
	}
	return nil
}

func (r *menuRepository) BatchUpdate(ctx context.Context, menus []*model.Menu) error {
	return r.query.WithContext(ctx).Menu.Save(menus...)
}

func (r *menuRepository) Delete(ctx context.Context, ids ...uint64) error {
	_, err := r.query.WithContext(ctx).Menu.Where(Menu.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *menuRepository) FindByID(ctx context.Context, id uint64) (*model.Menu, error) {
	return r.query.WithContext(ctx).Menu.Where(Menu.ID.Eq(id)).First()
}

func (r *menuRepository) List(ctx context.Context, query *model.MenuQuery) ([]*model.Menu, int64, error) {
	db := r.query.WithContext(ctx).Menu

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
		// db = db.Where("name LIKE ?", "%"+query.Name+"%")
		db = db.Where(Menu.Name.Like("%" + query.Name + "%"))
	}
	if query.Path != "" {
		// db = db.Where("path LIKE ?", "%"+query.Path+"%")
		db = db.Where(Menu.Path.Like("%" + query.Path + "%"))
	}
	if query.Component != "" {
		// db = db.Where("component LIKE ?", "%"+query.Component+"%")
		db = db.Where(Menu.Component.Like("%" + query.Component + "%"))
	}
	if query.Status != 0 {
		// db = db.Where("status = ?", query.Status)
		db = db.Where(Menu.Status.Eq(query.Status))
	}

	// 排序
	db.Order(Menu.Sort.Desc(), Menu.ID.Desc())

	// 统计总数
	total, err := db.Count()
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (query.Page - 1) * query.PageSize
	menus, err := db.Offset(offset).Limit(query.PageSize).Find()
	if err != nil {
		return nil, 0, err
	}

	return menus, total, nil
}

func (r *menuRepository) FindByParentID(ctx context.Context, parentID uint64) ([]*model.Menu, error) {
	menus, err := r.query.WithContext(ctx).Menu.Where(Menu.ParentID.Eq(parentID)).Find()
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) FindByRoleID(ctx context.Context, roleID uint64) ([]*model.Menu, error) {
	menus, err := r.query.WithContext(ctx).Menu.LeftJoin(RoleMenus, RoleMenus.MenuID.EqCol(Menu.ID)).Where(RoleMenus.RoleID.Eq(roleID)).Find()
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) FindByNames(ctx context.Context, names ...string) ([]*model.Menu, error) {
	menus, err := r.query.WithContext(ctx).Menu.Where(Menu.Name.In(names...)).Find()
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) FindByIDs(ctx context.Context, ids []uint64) ([]*model.Menu, error) {
	menus, err := r.query.WithContext(ctx).Menu.Where(Menu.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) FindAll(ctx context.Context) ([]*model.Menu, error) {
	menus, err := r.query.WithContext(ctx).Menu.Find()
	if err != nil {
		return nil, err
	}
	return menus, nil
}
