package repository

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/service"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
)

type roleMenuRepository struct {
	query *Query
}

func NewRoleMenuRepository(query *Query) service.RoleMenuRepository {
	return &roleMenuRepository{query: query}
}

func (r *roleMenuRepository) WithTx(tx *Query) service.RoleMenuRepository {
	return &roleMenuRepository{query: tx}
}

func (r *roleMenuRepository) Create(ctx context.Context, roleID, menuID uint64) error {
	return r.query.WithContext(ctx).RoleMenus.Create(&model.RoleMenus{
		RoleID: roleID,
		MenuID: menuID,
	})
}

func (r *roleMenuRepository) Delete(ctx context.Context, roleID, menuID uint64) error {
	_, err := r.query.WithContext(ctx).RoleMenus.Where(r.query.RoleMenus.RoleID.Eq(roleID),
		r.query.RoleMenus.MenuID.Eq(menuID)).Delete()
	return err
}

func (r *roleMenuRepository) DeleteByRoleID(ctx context.Context, roleID uint64) error {
	_, err := r.query.WithContext(ctx).RoleMenus.Where(r.query.RoleMenus.RoleID.Eq(roleID)).Delete()
	return err
}

func (r *roleMenuRepository) FindMenusByRoleID(ctx context.Context, roleID uint64) ([]*model.Menu, error) {
	menus, err := r.query.WithContext(ctx).Menu.LeftJoin(r.query.RoleMenus, r.query.RoleMenus.MenuID.EqCol(r.query.Menu.ID)).Where(r.query.RoleMenus.RoleID.Eq(roleID)).Find()
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *roleMenuRepository) FindRolesByMenuID(ctx context.Context, menuID uint64) ([]*model.Role, error) {
	roles, err := r.query.WithContext(ctx).Role.LeftJoin(r.query.RoleMenus, r.query.RoleMenus.RoleID.EqCol(r.query.Role.ID)).Where(r.query.RoleMenus.MenuID.Eq(menuID)).Find()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleMenuRepository) BatchCreate(ctx context.Context, roleID uint64, menuIDs []uint64) error {
	return r.query.Transaction(func(tx *Query) error {
		// 先删除原有的关联
		if _, err := tx.RoleMenus.Where(tx.RoleMenus.RoleID.Eq(roleID)).Delete(); err != nil {
			return err
		}

		// 批量创建新的关联
		var roleMenus []*model.RoleMenus
		for _, menuID := range menuIDs {
			roleMenus = append(roleMenus, &model.RoleMenus{
				RoleID: roleID,
				MenuID: menuID,
			})
		}
		return tx.RoleMenus.Create(roleMenus...)
	})
}

func (r *roleMenuRepository) FindMenusByRoleIDs(ctx context.Context, roleIDs ...uint64) ([]*model.Menu, error) {
	menus, err := r.query.WithContext(ctx).Menu.LeftJoin(r.query.RoleMenus, r.query.RoleMenus.MenuID.EqCol(r.query.Menu.ID)).Where(r.query.RoleMenus.RoleID.In(roleIDs...)).Find()
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *roleMenuRepository) FindRolesByMenuIDs(ctx context.Context, menuIDs []uint64) ([]*model.Role, error) {
	roles, err := r.query.WithContext(ctx).Role.LeftJoin(r.query.RoleMenus, r.query.RoleMenus.RoleID.EqCol(r.query.Role.ID)).Where(r.query.RoleMenus.MenuID.In(menuIDs...)).Find()
	if err != nil {
		return nil, err
	}
	return roles, nil
}
