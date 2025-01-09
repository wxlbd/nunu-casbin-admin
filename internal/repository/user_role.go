package repository

import (
	"context"
	"github.com/wxlbd/nunu-casbin-admin/internal/model"

	"gorm.io/gorm"
)

type UserRoleRepository interface {
	Create(ctx context.Context, userID, roleID uint64) error
	Delete(ctx context.Context, userID, roleID uint64) error
	FindRolesByUserID(ctx context.Context, userID uint64) ([]*model.Role, error)
	FindUsersByRoleID(ctx context.Context, roleID uint64) ([]*model.User, error)
}

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{
		db: db,
	}
}

func (r *userRoleRepository) Create(ctx context.Context, userID, roleID uint64) error {
	return r.db.WithContext(ctx).Create(map[string]interface{}{
		"user_id": userID,
		"role_id": roleID,
	}).Error
}

func (r *userRoleRepository) Delete(ctx context.Context, userID, roleID uint64) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Delete(&model.UserBelongsRole{}).Error
}

func (r *userRoleRepository) FindRolesByUserID(ctx context.Context, userID uint64) ([]*model.Role, error) {
	var roles []*model.Role
	err := r.db.WithContext(ctx).
		Joins("JOIN user_belongs_role ON user_belongs_role.role_id = role.id").
		Where("user_belongs_role.user_id = ?", userID).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *userRoleRepository) FindUsersByRoleID(ctx context.Context, roleID uint64) ([]*model.User, error) {
	var users []*model.User
	err := r.db.WithContext(ctx).
		Joins("JOIN user_belongs_role ON user_belongs_role.user_id = user.id").
		Where("user_belongs_role.role_id = ?", roleID).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
