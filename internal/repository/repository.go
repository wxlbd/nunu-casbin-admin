package repository

import (
	"gorm.io/gorm"
)

type Repository interface {
	User() UserRepository
	Role() RoleRepository
	Menu() MenuRepository
	UserRole() UserRoleRepository
	RoleMenu() RoleMenuRepository
	Transaction(fn func(Repository) error) error
	// DB 获取当前仓储使用的gorm.DB
	DB() *gorm.DB
}

type repository struct {
	query        *Query
	userRepo     UserRepository
	roleRepo     RoleRepository
	menuRepo     MenuRepository
	userRoleRepo UserRoleRepository
	roleMenuRepo RoleMenuRepository
}

func NewRepository(db *gorm.DB) Repository {
	SetDefault(db)
	return &repository{
		query:        Q,
		userRepo:     NewUserRepository(Q),
		roleRepo:     NewRoleRepository(Q),
		menuRepo:     NewMenuRepository(Q),
		userRoleRepo: NewUserRoleRepository(Q),
		roleMenuRepo: NewRoleMenuRepository(Q),
	}
}

func (r *repository) Transaction(fn func(Repository) error) error {
	return r.query.Transaction(func(tx *Query) error {
		return fn(r.clone(tx))
	})
}

func (r *repository) DB() *gorm.DB {
	return r.query.db
}

func (r *repository) clone(tx *Query) *repository {
	return &repository{
		query:        tx,
		userRepo:     NewUserRepository(tx),
		roleRepo:     NewRoleRepository(tx),
		menuRepo:     NewMenuRepository(tx),
		userRoleRepo: NewUserRoleRepository(tx),
		roleMenuRepo: NewRoleMenuRepository(tx),
	}
}

func (r *repository) Query() *Query {
	return r.query
}

func (r *repository) User() UserRepository {
	return r.userRepo
}

func (r *repository) Role() RoleRepository {
	return r.roleRepo
}

func (r *repository) Menu() MenuRepository {
	return r.menuRepo
}

func (r *repository) UserRole() UserRoleRepository {
	return r.userRoleRepo
}

func (r *repository) RoleMenu() RoleMenuRepository {
	return r.roleMenuRepo
}
