package repository

import (
	"github.com/wxlbd/gin-casbin-admin/internal/service"
	"gorm.io/gorm"
)

type repository struct {
	query        *Query
	userRepo     service.UserRepository
	roleRepo     service.RoleRepository
	menuRepo     service.MenuRepository
	userRoleRepo service.UserRoleRepository
	roleMenuRepo service.RoleMenuRepository
	dictTypeRepo service.DictTypeRepository
	dictDataRepo service.DictDataRepository
	db           *gorm.DB
}

func NewRepository(db *gorm.DB) service.Repository {
	SetDefault(db)
	return &repository{
		query:        Q,
		userRepo:     NewUserRepository(Q),
		roleRepo:     NewRoleRepository(Q),
		menuRepo:     NewMenuRepository(Q),
		userRoleRepo: NewUserRoleRepository(Q),
		roleMenuRepo: NewRoleMenuRepository(Q),
		dictTypeRepo: NewDictTypeRepository(Q),
		dictDataRepo: NewDictDataRepository(Q),
		db:           db,
	}
}

func (r *repository) Transaction(fn func(service.Repository) error) error {
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
		dictTypeRepo: NewDictTypeRepository(tx),
		dictDataRepo: NewDictDataRepository(tx),
		db:           r.db,
	}
}

func (r *repository) Query() *Query {
	return r.query
}

func (r *repository) User() service.UserRepository {
	return r.userRepo
}

func (r *repository) Role() service.RoleRepository {
	return r.roleRepo
}

func (r *repository) Menu() service.MenuRepository {
	return r.menuRepo
}

func (r *repository) UserRole() service.UserRoleRepository {
	return r.userRoleRepo
}

func (r *repository) RoleMenu() service.RoleMenuRepository {
	return r.roleMenuRepo
}

func (r *repository) DictType() service.DictTypeRepository {
	return r.dictTypeRepo
}

func (r *repository) DictData() service.DictDataRepository {
	return r.dictDataRepo
}

func (r *repository) SysMenu() service.SysMenuRepository {
	return NewSysMenuRepository(r.query)
}
