package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/wxlbd/gin-casbin-admin/internal/handler"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/jwtx"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
)

type service struct {
	repo    Repository
	userSvc handler.UserService
	roleSvc handler.RoleService
	menuSvc handler.MenuService
	dict    handler.DictService
	jwt     *jwtx.JWT
}

func NewService(logger *log.Logger, repo Repository, cfg *config.Config, rdb *redis.Client, j *jwtx.JWT, enforcer *casbin.Enforcer) (handler.Service, error) {
	svc := &service{
		repo: repo,
		jwt:  j,
	}

	svc.userSvc = NewUserService(logger, repo, j)
	svc.roleSvc = NewRoleService(repo, enforcer)
	svc.menuSvc = NewMenuService(repo, enforcer)
	svc.dict = NewDictService(logger, repo)

	return svc, nil
}

func (s *service) User() handler.UserService {
	return s.userSvc
}

func (s *service) Role() handler.RoleService {
	return s.roleSvc
}

func (s *service) Menu() handler.MenuService {
	return s.menuSvc
}

func (s *service) Dict() handler.DictService {
	return s.dict
}
