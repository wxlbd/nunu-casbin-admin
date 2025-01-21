package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/jwtx"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
)

type Service interface {
	User() UserService
	Role() RoleService
	Menu() MenuService
	Dict() DictService
}

type service struct {
	repo    Repository
	userSvc UserService
	roleSvc RoleService
	menuSvc MenuService
	dict    DictService
	jwt     *jwtx.JWT
}

func NewService(logger *log.Logger, repo Repository, cfg *config.Config, rdb *redis.Client, j *jwtx.JWT, enforcer *casbin.Enforcer) (Service, error) {
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

func (s *service) User() UserService {
	return s.userSvc
}

func (s *service) Role() RoleService {
	return s.roleSvc
}

func (s *service) Menu() MenuService {
	return s.menuSvc
}

func (s *service) Dict() DictService {
	return s.dict
}
