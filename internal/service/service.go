package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/wxlbd/nunu-casbin-admin/internal/repository"
	"github.com/wxlbd/nunu-casbin-admin/pkg/config"
	"github.com/wxlbd/nunu-casbin-admin/pkg/jwt"
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
	jwt     *jwt.JWT
}

func NewService(repo repository.Repository, cfg *config.Config, rdb *redis.Client) (Service, error) {

	j := jwt.New(jwt.Config{
		AccessSecret:  cfg.JWT.AccessSecret,
		RefreshSecret: cfg.JWT.RefreshSecret,
		AccessExpire:  cfg.JWT.AccessExpire,
		RefreshExpire: cfg.JWT.RefreshExpire,
		Issuer:        cfg.JWT.Issuer,
	}, rdb)

	svc := &service{
		repo: repo,
		jwt:  j,
	}

	svc.userSvc = NewUserService(repo, j)
	svc.roleSvc = NewRoleService(repo)
	svc.menuSvc = NewMenuService(repo)

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
