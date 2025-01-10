package dto

import (
	"github.com/wxlbd/nunu-casbin-admin/internal/handler/response"
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"time"
)

// ToUserResponse 将 User 模型转换为响应 DTO
func ToUserResponse(user *model.User) *response.UserResponse {
	if user == nil {
		return nil
	}
	return &response.UserResponse{
		ID:             int(user.ID),
		Username:       user.Username,
		Nickname:       user.Nickname,
		Phone:          user.Phone,
		Email:          user.Email,
		Avatar:         user.Avatar,
		Status:         int(user.Status),
		LoginTime:      user.LoginTime.Format(time.DateTime),
		CreatedBy:      int(user.CreatedBy),
		UpdatedBy:      int(user.UpdatedBy),
		CreatedAt:      user.CreatedAt.Format(time.DateTime),
		UpdatedAt:      user.UpdatedAt.Format(time.DateTime),
		Remark:         user.Remark,
		UserType:       user.UserType,
		Signed:         user.Signed,
		LoginIp:        user.LoginIp,
		BackendSetting: user.BackendSetting,
	}
}

// ToUserResponseList 将用户列表转换为响应 DTO 列表
func ToUserResponseList(users []*model.User) []*response.UserResponse {
	if users == nil {
		return nil
	}
	list := make([]*response.UserResponse, 0, len(users))
	for _, user := range users {
		list = append(list, ToUserResponse(user))
	}
	return list
}

// ToUserListResponse 转换为分页响应
func ToUserListResponse(users []*model.User, total int64) *response.UserListResponse {
	return &response.UserListResponse{
		List:  ToUserResponseList(users),
		Total: total,
	}
}
