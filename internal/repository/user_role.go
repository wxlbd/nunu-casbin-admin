package repository

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
)

type UserRoleRepository interface {
	WithTx(tx *Query) UserRoleRepository
	Create(ctx context.Context, userRoles ...*model.UserRoles) error
	DeleteByUserID(ctx context.Context, userID uint64) error
	FindRolesByUserID(ctx context.Context, userID uint64) ([]*model.Role, error)
}

type userRoleRepository struct {
	query *Query
}

func NewUserRoleRepository(query *Query) UserRoleRepository {
	return &userRoleRepository{query: query}
}

func (r *userRoleRepository) WithTx(tx *Query) UserRoleRepository {
	return &userRoleRepository{query: tx}
}

func (r *userRoleRepository) DeleteByUserID(ctx context.Context, userID uint64) error {
	_, err := r.query.WithContext(ctx).UserRoles.Where(r.query.UserRoles.UserID.Eq(userID)).Delete()
	return err
}

func (r *userRoleRepository) Create(ctx context.Context, userRoles ...*model.UserRoles) error {
	return r.query.WithContext(ctx).UserRoles.Create(userRoles...)
}

func (r *userRoleRepository) FindRolesByUserID(ctx context.Context, userID uint64) ([]*model.Role, error) {
	roles, err := r.query.WithContext(ctx).Role.
		LeftJoin(r.query.UserRoles, r.query.UserRoles.RoleID.EqCol(r.query.Role.ID)).
		Where(r.query.UserRoles.UserID.Eq(userID)).
		Find()
	if err != nil {
		return nil, err
	}
	return roles, nil
}
