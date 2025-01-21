package repository

import (
	"context"
	"errors"

	"github.com/wxlbd/gin-casbin-admin/internal/service"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"gorm.io/gorm"
)

type roleRepository struct {
	query *Query
}

func (r *roleRepository) FindByCodes(ctx context.Context, codes ...string) ([]*model.Role, error) {
	roles, err := r.query.WithContext(ctx).Role.Where(r.query.Role.Code.In(codes...)).Find()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// FindByIDs implements RoleRepository.
func (r *roleRepository) FindByIDs(ctx context.Context, ids []uint64) ([]*model.Role, error) {
	roles, err := r.query.WithContext(ctx).Role.Where(r.query.Role.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func NewRoleRepository(query *Query) service.RoleRepository {
	return &roleRepository{query: query}
}

func (r *roleRepository) WithTx(tx *Query) service.RoleRepository {
	return &roleRepository{query: tx}
}

func (r *roleRepository) Create(ctx context.Context, role *model.Role) error {
	return r.query.WithContext(ctx).Role.Create(role)
}

func (r *roleRepository) Update(ctx context.Context, role *model.Role) error {
	_, err := r.query.WithContext(ctx).Role.Updates(role)
	return err
}

func (r *roleRepository) Delete(ctx context.Context, ids ...uint64) error {
	_, err := r.query.WithContext(ctx).Role.Where(r.query.Role.ID.In(ids...)).Delete()
	return err
}

func (r *roleRepository) FindByID(ctx context.Context, id uint64) (*model.Role, error) {
	role, err := r.query.WithContext(ctx).Role.Where(r.query.Role.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) FindByCode(ctx context.Context, code string) (*model.Role, error) {
	role, err := r.query.WithContext(ctx).Role.Where(r.query.Role.Code.Eq(code)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) List(ctx context.Context, query *model.RoleQuery) ([]*model.Role, int64, error) {
	db := r.query.WithContext(ctx).Role
	if query.Name != "" {
		db = db.Where(r.query.Role.Name.Like("%" + query.Name + "%"))
	}
	if query.Code != "" {
		db = db.Where(r.query.Role.Code.Like("%" + query.Code + "%"))
	}
	if query.Status != 0 {
		db = db.Where(r.query.Role.Status.Eq(query.Status))
	}
	total, err := db.Count()
	if err != nil {
		return nil, 0, err
	}
	roles, err := db.Offset(query.GetOffset()).Limit(query.PageSize).Find()
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}
