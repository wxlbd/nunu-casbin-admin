package service

import (
	"context"

	"gorm.io/gorm"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
)

type DictTypeRepository interface {
	Create(ctx context.Context, dict *model.DictType) error
	Update(ctx context.Context, dict *model.DictType) error
	Delete(ctx context.Context, ids ...int64) error
	FindByID(ctx context.Context, id int64) (*model.DictType, error)
	FindByCode(ctx context.Context, code string) (*model.DictType, error)
	List(ctx context.Context, query *model.DictTypeQuery) ([]*model.DictType, int64, error)
}

type DictDataRepository interface {
	Create(ctx context.Context, data *model.DictDatum) error
	Update(ctx context.Context, data *model.DictDatum) error
	Delete(ctx context.Context, ids ...int64) error
	FindByID(ctx context.Context, id int64) (*model.DictDatum, error)
	FindByTypeCode(ctx context.Context, typeCode string) ([]*model.DictDatum, error)
	List(ctx context.Context, query *model.DictDataQuery) ([]*model.DictDatum, int64, error)
}

type MenuRepository interface {
	Create(ctx context.Context, menu *model.Menu) error
	BatchCreate(ctx context.Context, menus []*model.Menu) error
	Update(ctx context.Context, menu *model.Menu) error
	BatchUpdate(ctx context.Context, menus []*model.Menu) error
	Delete(ctx context.Context, ids ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Menu, error)
	FindByIDs(ctx context.Context, ids []uint64) ([]*model.Menu, error)
	FindByNames(ctx context.Context, names ...string) ([]*model.Menu, error)
	FindAll(ctx context.Context) ([]*model.Menu, error)
	List(ctx context.Context, query *model.MenuQuery) ([]*model.Menu, int64, error)
	FindByParentID(ctx context.Context, parentID uint64) ([]*model.Menu, error)
	FindByRoleID(ctx context.Context, roleID uint64) ([]*model.Menu, error)
}

type RoleRepository interface {
	Create(ctx context.Context, role *model.Role) error
	Update(ctx context.Context, role *model.Role) error
	Delete(ctx context.Context, ids ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Role, error)
	FindByCode(ctx context.Context, code string) (*model.Role, error)
	List(ctx context.Context, query *model.RoleQuery) ([]*model.Role, int64, error)
	// FindByIDs 根据角色ID列表查询角色
	FindByIDs(ctx context.Context, ids []uint64) ([]*model.Role, error)
	FindByCodes(ctx context.Context, codes ...string) ([]*model.Role, error)
}

type RoleMenuRepository interface {
	Create(ctx context.Context, roleID, menuID uint64) error
	Delete(ctx context.Context, roleID, menuID uint64) error
	DeleteByRoleID(ctx context.Context, roleID uint64) error
	FindMenusByRoleID(ctx context.Context, roleID uint64) ([]*model.Menu, error)
	FindRolesByMenuID(ctx context.Context, menuID uint64) ([]*model.Role, error)
	BatchCreate(ctx context.Context, roleID uint64, menuIDs []uint64) error
	FindMenusByRoleIDs(ctx context.Context, roleIDs ...uint64) ([]*model.Menu, error)
	FindRolesByMenuIDs(ctx context.Context, menuIDs []uint64) ([]*model.Role, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, ids ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context, query *model.UserQuery) ([]*model.User, int64, error)
}

type UserRoleRepository interface {
	Create(ctx context.Context, userRoles ...*model.UserRoles) error
	DeleteByUserID(ctx context.Context, userID uint64) error
	FindRolesByUserID(ctx context.Context, userID uint64) ([]*model.Role, error)
}

type Repository interface {
	User() UserRepository
	Role() RoleRepository
	Menu() MenuRepository
	UserRole() UserRoleRepository
	RoleMenu() RoleMenuRepository
	DictType() DictTypeRepository
	DictData() DictDataRepository
	Transaction(fn func(Repository) error) error
	// DB 获取当前仓储使用的gorm.DB
	DB() *gorm.DB
}