package repository

import "gorm.io/gorm"

type Repository interface {
	User() UserRepository
	Role() RoleRepository
	Menu() MenuRepository
	UserRole() UserRoleRepository
	RoleMenu() RoleMenuRepository
	DB() *gorm.DB
}

type repository struct {
	db           *gorm.DB
	userRepo     UserRepository
	roleRepo     RoleRepository
	menuRepo     MenuRepository
	userRoleRepo UserRoleRepository
	roleMenuRepo RoleMenuRepository
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db:           db,
		userRepo:     NewUserRepository(db),
		roleRepo:     NewRoleRepository(db),
		menuRepo:     NewMenuRepository(db),
		userRoleRepo: NewUserRoleRepository(db),
		roleMenuRepo: NewRoleMenuRepository(db),
	}
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

func (r *repository) DB() *gorm.DB {
	return r.db
}
