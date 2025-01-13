package service

import (
	"context"
	"errors"

	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"github.com/wxlbd/nunu-casbin-admin/internal/repository"
	"github.com/wxlbd/nunu-casbin-admin/pkg/jwtx"

	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, ids ...uint64) error
	FindByID(ctx context.Context, id uint64) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context, query *model.UserQuery) ([]*model.User, int64, error)
	UpdatePassword(ctx context.Context, id uint64, oldPassword, newPassword string) error
	AssignRoles(ctx context.Context, userID uint64, roleIDs []uint64) error
	Login(ctx context.Context, username, password string) (accessToken, refreshToken string, err error)
	RefreshToken(ctx context.Context, refreshToken string) (newAccessToken, newRefreshToken string, err error)
	Logout(ctx context.Context, token string) error
	GetUserRoles(ctx context.Context, userID uint64) ([]*model.Role, error)
}

type userService struct {
	repo repository.Repository
	jwt  *jwtx.JWT
}

func NewUserService(repo repository.Repository, jwt *jwtx.JWT) UserService {
	return &userService{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *userService) Create(ctx context.Context, user *model.User) error {
	// 检查用户名是否已存在
	existUser, _ := s.repo.User().FindByUsername(ctx, user.Username)
	if existUser != nil {
		return errors.New("用户名已存在")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.LoginTime = time.Unix(0, 0)
	return s.repo.User().Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, user *model.User) error {
	existUser, err := s.repo.User().FindByID(ctx, user.ID)
	if err != nil {
		return err
	}
	if existUser == nil {
		return errors.New("用户不存在")
	}

	// 如果修改了用户名，需要检查新用户名是否已存在
	if user.Username != existUser.Username {
		if exist, _ := s.repo.User().FindByUsername(ctx, user.Username); exist != nil {
			return errors.New("用户名已存在")
		}
	}

	return s.repo.User().Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, ids ...uint64) error {
	return s.repo.User().Delete(ctx, ids...)
}

func (s *userService) FindByID(ctx context.Context, id uint64) (*model.User, error) {
	return s.repo.User().FindByID(ctx, id)
}

func (s *userService) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.repo.User().FindByUsername(ctx, username)
}

func (s *userService) List(ctx context.Context, query *model.UserQuery) ([]*model.User, int64, error) {
	return s.repo.User().List(ctx, query)
}

func (s *userService) UpdatePassword(ctx context.Context, id uint64, oldPassword, newPassword string) error {
	user, err := s.repo.User().FindByID(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("用户不存在")
	}

	// 验证旧密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("旧密码错误")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.repo.User().Update(ctx, user)
}

func (s *userService) AssignRoles(ctx context.Context, userID uint64, roleIDs []uint64) error {
	// 开启事务
	// 先删除用户所有角色，再重新分配
	for _, roleID := range roleIDs {
		if err := s.repo.UserRole().Create(ctx, userID, roleID); err != nil {
			return err
		}
	}
	return nil
}

func (s *userService) Login(ctx context.Context, username, password string) (accessToken, refreshToken string, err error) {
	user, err := s.repo.User().FindByUsername(ctx, username)
	if err != nil {
		return "", "", err
	}
	if user == nil {
		return "", "", errors.New("用户不存在")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("密码错误")
	}

	// 生成 token
	accessToken, refreshToken, err = s.jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", "", err
	}

	// 更新登录信息
	user.LoginTime = time.Now()
	if err := s.repo.User().Update(ctx, user); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *userService) RefreshToken(ctx context.Context, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	return s.jwt.RefreshToken(ctx, refreshToken)
}

func (s *userService) Logout(ctx context.Context, token string) error {
	// 解析 token
	claims, err := s.jwt.ParseToken(ctx, token, false)
	if err != nil {
		return err
	}

	// 将 token 加入黑名单
	return s.jwt.AddToBlacklist(ctx, token, claims)
}

func (s *userService) GetUserRoles(ctx context.Context, userID uint64) ([]*model.Role, error) {
	// 检查用户是否存在
	user, err := s.repo.User().FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	// 获取用户的角色列表
	return s.repo.UserRole().FindRolesByUserID(ctx, userID)
}
