package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/wxlbd/nunu-casbin-admin/internal/repository"
	"github.com/wxlbd/nunu-casbin-admin/pkg/config"
	"github.com/wxlbd/nunu-casbin-admin/pkg/jwtx"
	"github.com/wxlbd/nunu-casbin-admin/pkg/log"
)

type Service interface {
	User() UserService
	Role() RoleService
	Menu() MenuService
}

type service struct {
	repo    repository.Repository
	userSvc UserService
	roleSvc RoleService
	menuSvc MenuService
	jwt     *jwtx.JWT
}

func NewService(logger *log.Logger, repo repository.Repository, cfg *config.Config, rdb *redis.Client, j *jwtx.JWT, enforcer *casbin.Enforcer) (Service, error) {
	svc := &service{
		repo: repo,
		jwt:  j,
	}

	svc.userSvc = NewUserService(logger, repo, j)
	svc.roleSvc = NewRoleService(repo, enforcer)
	svc.menuSvc = NewMenuService(repo, enforcer)

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
