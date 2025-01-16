package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/wxlbd/gin-casbin-admin/pkg/errors"

	"github.com/casbin/casbin/v2"
	"github.com/wxlbd/gin-casbin-admin/internal/dto"
	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/internal/repository"
	"gorm.io/gorm"
)

type RoleService interface {
	Create(ctx context.Context, role *model.Role) error
	Update(ctx context.Context, role *model.Role) error
	Delete(ctx context.Context, id ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.Role, error)
	List(ctx context.Context, req *dto.RoleListRequest) ([]*model.Role, int64, error)
	AssignMenus(ctx context.Context, roleID uint64, menus []string) error
	GetRoleMenus(ctx context.Context, roleID uint64) ([]*model.Menu, error)
}

type roleService struct {
	repo     repository.Repository
	enforcer *casbin.Enforcer
}

func NewRoleService(repo repository.Repository, enforcer *casbin.Enforcer) RoleService {
	return &roleService{
		repo:     repo,
		enforcer: enforcer,
	}
}

func (s *roleService) Create(ctx context.Context, role *model.Role) error {
	if s.IsCodeExists(ctx, role.Code) {
		return errors.WithMsg(errors.AlreadyExists, "角色代码已存在")
	}
	return s.repo.Role().Create(ctx, role)
}

func (s *roleService) IsCodeExists(ctx context.Context, code string) bool {
	roles, _ := s.repo.Role().FindByCodes(ctx, code)
	return len(roles) > 0
}

func (s *roleService) Update(ctx context.Context, role *model.Role) error {
	existRole, err := s.repo.Role().FindByID(ctx, role.ID)
	if err != nil {
		return err
	}
	if existRole == nil {
		return errors.WithMsg(errors.NotFound, "角色不存在")
	}

	// 如果修改了角色代码，需要检查新代码是否已存在
	if role.Code != existRole.Code {
		if s.IsCodeExists(ctx, role.Code) {
			return errors.WithMsg(errors.AlreadyExists, "角色代码已存在")
		}
	}

	return s.repo.Role().Update(ctx, role)
}

func (s *roleService) Delete(ctx context.Context, ids ...uint64) error {
	roles, err := s.repo.Role().FindByIDs(ctx, ids)
	if err != nil {
		return err
	}
	if len(roles) == 0 {
		return errors.WithMsg(errors.NotFound, "角色不存在")
	}
	// 删除该角色的所有权限策略
	for _, role := range roles {
		_, err = s.enforcer.DeletePermissionsForUser(role.Code)
		if err != nil {
			return err
		}
	}
	return s.repo.Role().Delete(ctx, ids...)
}

func (s *roleService) FindByID(ctx context.Context, id uint64) (*model.Role, error) {
	return s.repo.Role().FindByID(ctx, id)
}

func (s *roleService) List(ctx context.Context, req *dto.RoleListRequest) ([]*model.Role, int64, error) {
	query := req.ToModel()
	return s.repo.Role().List(ctx, query)
}

func (s *roleService) AssignMenus(ctx context.Context, roleID uint64, names []string) error {
	// 开启事务
	return s.repo.DB().Transaction(func(tx *gorm.DB) error {
		// 1. 先删除原有的角色-菜单关联
		if err := s.repo.RoleMenu().DeleteByRoleID(ctx, roleID); err != nil {
			return err
		}

		// 2. 获取所有按钮类型的菜单（即 API）
		menus, err := s.repo.Menu().FindByName(ctx, names...)
		if err != nil {
			return err
		}

		// 3. 更新 Casbin 策略
		role, err := s.repo.Role().FindByID(ctx, roleID)
		if err != nil {
			return err
		}

		// 先删除该角色的所有策略
		_, err = s.enforcer.DeletePermissionsForUser(role.Code)
		if err != nil {
			return err
		}

		var menuIDs []uint64
		// 添加新的策略
		for _, menu := range menus {
			menuIDs = append(menuIDs, menu.ID)
			if menu.Meta.Type == "B" { // 按钮类型
				// 根据菜单名称生成 API 路径和方法
				// 例如：permission:user:save -> POST /api/user
				path, method := convertMenuToAPI(menu.Name)
				_, err = s.enforcer.AddPolicy(role.Code, path, method)
				if err != nil {
					return err
				}
			}
		}

		// 4. 创建新的角色-菜单关联
		return s.repo.RoleMenu().BatchCreate(ctx, roleID, menuIDs)
	})
}

// convertMenuToAPI 将菜单名称转换为 API 路径和方法
// 示例: system:role:get:menus -> GET /api/system/role/:id/menus
func convertMenuToAPI(menuName string) (path, method string) {
	// 基础路径前缀
	const apiPrefix = "/api"

	// 分割菜单名称
	parts := strings.Split(menuName, ":")
	if len(parts) < 3 {
		return "", ""
	}

	// 获取模块、资源和动作
	module := parts[0]   // 例如: system
	resource := parts[1] // 例如: role

	// 处理复合动作（如 get:menus）
	var action, subResource string
	if len(parts) >= 4 {
		action = parts[2]      // 例如: get
		subResource = parts[3] // 例如: menus
	} else {
		action = parts[2] // 例如: create
	}

	// 动作映射表
	actionMap := map[string]struct {
		method     string
		pathSuffix string
	}{
		"create": {"POST", ""},
		"save":   {"POST", ""},
		"update": {"PUT", ":id"},
		"delete": {"DELETE", ":ids"},
		"get":    {"GET", ":id"},
		"detail": {"GET", ":id"},
		"list":   {"GET", ""},
		"index":  {"GET", ""},

		// 扩展的业务操作
		"enable":  {"PATCH", "enable"},
		"disable": {"PATCH", "disable"},
		"assign":  {"POST", "assign"},
		"revoke":  {"POST", "revoke"},
		"upload":  {"POST", "upload"},
		"export":  {"GET", "export"},
		"import":  {"POST", "import"},
		"batch":   {"POST", "batch"},
		"tree":    {"GET", "tree"},
		"status":  {"PATCH", "status"},
		"set":     {"PUT", ":id"},
	}

	// 获取 HTTP 方法
	item, ok := actionMap[action]
	if !ok {
		return "", ""
	}

	// 构建路径
	if subResource != "" {
		// 对于子资源路径：/api/system/role/:id/menus
		path = fmt.Sprintf("%s/%s/%s/%s/%s",
			apiPrefix,
			module,
			resource,
			item.pathSuffix,
			subResource,
		)
	} else {
		// 对于普通路径：/api/system/role
		path = fmt.Sprintf("%s/%s/%s",
			apiPrefix,
			module,
			resource,
		)
		if item.pathSuffix != "" {
			path = fmt.Sprintf("%s/%s", path, item.pathSuffix)
		}
	}
	return path, item.method
}

func (s *roleService) GetRoleMenus(ctx context.Context, roleID uint64) ([]*model.Menu, error) {
	// 检查角色是否存在
	role, err := s.repo.Role().FindByID(ctx, roleID)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.WithMsg(errors.NotFound, "角色不存在")
	}

	// 获取角色的菜单列表
	return s.repo.RoleMenu().FindMenusByRoleID(ctx, roleID)
}
