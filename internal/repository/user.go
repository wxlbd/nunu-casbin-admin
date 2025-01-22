package repository

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/service"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
)

type userRepository struct {
	query *Query
}

func NewUserRepository(query *Query) service.UserRepository {
	return &userRepository{
		query: query,
	}
}

func (r *userRepository) WithTx(tx *Query) service.UserRepository {
	return &userRepository{
		query: tx,
	}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.query.WithContext(ctx).User.Create(user)
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.query.WithContext(ctx).User.Save(user)
}

func (r *userRepository) Delete(ctx context.Context, ids ...uint64) error {
	_, err := r.query.WithContext(ctx).User.Where(r.query.User.ID.In(ids...)).Delete()
	return err
}

func (r *userRepository) FindByID(ctx context.Context, id uint64) (*model.User, error) {
	return r.query.WithContext(ctx).User.Where(r.query.User.ID.Eq(id)).First()
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := r.query.WithContext(ctx).User.Where(r.query.User.Username.Eq(username)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) List(ctx context.Context, query *model.UserQuery) ([]*model.User, int64, error) {
	q := r.query.WithContext(ctx).User
	// 构建查询条件
	if query != nil {
		if query.Username != "" {
			q = q.Where(r.query.User.Username.Like("%" + query.Username + "%"))
		}
		if query.Nickname != "" {
			q = q.Where(r.query.User.Nickname.Like("%" + query.Nickname + "%"))
		}
		if query.Status != 0 {
			q = q.Where(r.query.User.Status.Eq(query.Status))
		}
	}

	// 统计总数
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (query.Page - 1) * query.PageSize
	users, err := q.Offset(offset).Limit(query.PageSize).Find()
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
