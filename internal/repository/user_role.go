package repository

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/service"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
)

type userRoleRepository struct {
	query *Query
}

func NewUserRoleRepository(query *Query) service.UserRoleRepository {
	return &userRoleRepository{query: query}
}

func (r *userRoleRepository) WithTx(tx *Query) service.UserRoleRepository {
	return &userRoleRepository{query: tx}
}

func (r *userRoleRepository) DeleteByUserID(ctx context.Context, userID uint64) error {
	_, err := r.query.WithContext(ctx).UserRoles.Where(r.query.UserRoles.UserID.Eq(userID)).Delete()
	return err
}

func (r *userRoleRepository) Create(ctx context.Context, userRoles ...*model.UserRoles) error {
	return r.query.WithContext(ctx).UserRoles.Create(userRoles...)
}

// FindRolesByUserID 根据用户ID查找角色。
// 该方法通过查询数据库，获取与特定用户ID关联的所有角色。
// 参数:
//
//	ctx - 上下文，用于处理请求和传递请求范围内的值。
//	userID - 用户的唯一标识符，用于查找该用户的角色。
//
// 返回值:
//
//	[]*model.Role - 一个角色指针的切片，包含找到的角色。
//	error - 如果查询过程中发生错误，返回该错误。
func (r *userRoleRepository) FindRolesByUserID(ctx context.Context, userID uint64) ([]*model.Role, error) {
	// 执行数据库查询，通过用户ID查找对应的角色。
	roles, err := r.query.WithContext(ctx).Role.
		LeftJoin(r.query.UserRoles, r.query.UserRoles.RoleID.EqCol(r.query.Role.ID)).
		Where(r.query.UserRoles.UserID.Eq(userID)).
		Find()
	if err != nil {
		// 如果查询过程中发生错误，返回nil和错误。
		return nil, err
	}
	// 返回找到的角色和nil作为错误。
	return roles, nil
}
