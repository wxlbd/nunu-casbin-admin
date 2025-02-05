package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/internal/service"
)

type sysMenuRepository struct {
	*Query
}

func (r *sysMenuRepository) FindByIDs(ctx context.Context, ids ...uint64) ([]*model.SysMenu, error) {
	var i64s = make([]int64, len(ids))
	for i, id := range ids {
		i64s[i] = int64(id)
	}
	return r.WithContext(ctx).SysMenu.Where(r.SysMenu.ID.In(i64s...)).Find()
}

func NewSysMenuRepository(query *Query) service.SysMenuRepository {
	return &sysMenuRepository{
		Query: query,
	}
}

func (r *sysMenuRepository) Create(ctx context.Context, menu *model.SysMenu) error {
	return r.WithContext(ctx).SysMenu.Create(menu)
}

func (r *sysMenuRepository) Save(ctx context.Context, menu *model.SysMenu) error {
	return r.WithContext(ctx).SysMenu.Save(menu)
}

func (r *sysMenuRepository) Delete(ctx context.Context, ids ...int64) error {
	_, err := r.WithContext(ctx).SysMenu.Where(r.SysMenu.ID.In(ids...)).Delete()
	return err
}

func (r *sysMenuRepository) Get(ctx context.Context, id int64) (*model.SysMenu, error) {
	return r.WithContext(ctx).SysMenu.Where(r.SysMenu.ID.Eq(id)).First()
}

func (r *sysMenuRepository) FindByTitle(ctx context.Context, title string) (*model.SysMenu, error) {

	menu, err := r.WithContext(ctx).SysMenu.Where(r.SysMenu.Title.Eq(title)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 或者返回一个自定义错误
		}
		return nil, err
	}
	return menu, nil
}

func (r *sysMenuRepository) FindByParentID(ctx context.Context, parentID int64) ([]*model.SysMenu, error) {
	return r.WithContext(ctx).SysMenu.Where(r.SysMenu.ParentID.Eq(parentID)).Find()
}

func (r *sysMenuRepository) List(ctx context.Context, query *model.SysMenuQuery) ([]*model.SysMenu, int64, error) {
	q := r.WithContext(ctx).SysMenu

	// 构建查询条件
	if query.Title != "" {
		q = q.Where(r.SysMenu.Title.Like("%" + query.Title + "%"))
	}
	if query.Status != 0 {
		q = q.Where(r.SysMenu.Status.Eq(query.Status))
	}
	if query.MenuType != 0 {
		q = q.Where(r.SysMenu.MenuType.Eq(query.MenuType))
	}

	// 获取总数
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (query.Page - 1) * query.PageSize
	list, err := q.Order(r.SysMenu.Rank).
		Offset(offset).
		Limit(query.PageSize).
		Find()

	return list, total, err
}

// FindAll 获取所有菜单
func (r *sysMenuRepository) FindAll(ctx context.Context) ([]*model.SysMenu, error) {
	return r.WithContext(ctx).SysMenu.Order(r.SysMenu.Rank).Find()
}

// FindByRoleIDs 根据角色ID列表获取菜单
func (r *sysMenuRepository) FindByRoleIDs(ctx context.Context, roleIDs ...uint64) ([]*model.SysMenu, error) {
	// 1. 先查询角色菜单关系表获取菜单ID
	var menuIDs []int64
	err := r.WithContext(ctx).RoleMenus.
		Select(r.RoleMenus.MenuID).
		Where(r.RoleMenus.RoleID.In(roleIDs...)).
		Scan(&menuIDs)
	if err != nil {
		return nil, err
	}

	if len(menuIDs) == 0 {
		return nil, nil
	}

	// 2. 根据菜单ID获取菜单信息
	return r.WithContext(ctx).SysMenu.
		Where(r.SysMenu.ID.In(menuIDs...)).
		Order(r.SysMenu.Rank).
		Find()
}
