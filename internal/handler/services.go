package handler

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/dto"
	"github.com/wxlbd/gin-casbin-admin/internal/model"
)

type MenuService interface {
	Create(ctx context.Context, req *dto.CreateMenuRequest) error
	Update(ctx context.Context, req *dto.UpdateMenuRequest) error
	Delete(ctx context.Context, id ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Menu, error)
	List(ctx context.Context, query *model.MenuQuery) ([]*model.Menu, int64, error)
	GetMenuTree(ctx context.Context) ([]*model.MenuTree, error)
	GetUserMenus(ctx context.Context, userID uint64) ([]*model.MenuTree, error)
	GetAllMenus(ctx context.Context) ([]*model.Menu, error)
}

type DictService interface {
	// CreateDictType DictType
	CreateDictType(ctx context.Context, req *dto.DictTypeRequest) error
	UpdateDictType(ctx context.Context, req *dto.DictTypeRequest) error
	DeleteDictType(ctx context.Context, ids ...int64) error
	GetDictType(ctx context.Context, id int64) (*model.DictType, error)
	ListDictType(ctx context.Context, query *model.DictTypeQuery) ([]*model.DictType, int64, error)

	// CreateDictData DictData
	CreateDictData(ctx context.Context, req *dto.DictDataRequest) error
	UpdateDictData(ctx context.Context, req *dto.DictDataRequest) error
	DeleteDictData(ctx context.Context, ids ...int64) error
	GetDictData(ctx context.Context, id int64) (*model.DictDatum, error)
	ListDictData(ctx context.Context, query *model.DictDataQuery) ([]*model.DictDatum, int64, error)
	GetDictDataByType(ctx context.Context, typeCode string) ([]*model.DictDatum, error)
}

type RoleService interface {
	Create(ctx context.Context, role *model.Role) error
	Update(ctx context.Context, role *model.Role) error
	Delete(ctx context.Context, id ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Role, error)
	List(ctx context.Context, req *dto.RoleListRequest) ([]*model.Role, int64, error)
	GetRoleMenus(ctx context.Context, roleID uint64) ([]*model.Menu, error)
	AssignMenuByIds(ctx context.Context, roleID uint64, menuIds []uint64) error
	// GetAllRoles 获取所有角色
	GetAllRoles(ctx context.Context) ([]*model.Role, error)
}

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, ids ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context, query *model.UserQuery) ([]*model.User, int64, error)
	UpdatePassword(ctx context.Context, id uint64, oldPassword, newPassword string) error
	AssignRoles(ctx context.Context, userID uint64, roleIds []uint64) error
	Login(ctx context.Context, username, password string) (accessToken, refreshToken string, err error)
	RefreshToken(ctx context.Context, refreshToken string) (newAccessToken, newRefreshToken string, err error)
	Logout(ctx context.Context, token string) error
	GetUserRoles(ctx context.Context, userID uint64) ([]*model.Role, error)
}

type CaptchaService interface {
	Generate(ctx context.Context) (id, b64s string, err error)
	Verify(ctx context.Context, id, answer string) bool
}

type SysMenuService interface {
	Create(ctx context.Context, menu *model.SysMenu) error
	Update(ctx context.Context, menu *model.SysMenu) error
	Delete(ctx context.Context, ids ...int64) error
	List(ctx context.Context, query *model.SysMenuQuery) ([]*model.SysMenu, int64, error)
	GetMenuTree(ctx context.Context) ([]*model.SysMenuTree, error)
	GetUserMenuTree(ctx context.Context, userID uint64) ([]*model.SysMenuTree, error)
	GetAllMenus(ctx context.Context) ([]*model.SysMenu, error)
}

type Service interface {
	User() UserService
	Role() RoleService
	Menu() MenuService
	Dict() DictService
	Captcha() CaptchaService
	SysMenu() SysMenuService
}
