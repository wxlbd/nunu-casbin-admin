package dto

import (
	"github.com/wxlbd/nunu-casbin-admin/internal/handler/response"
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"time"
)

func ToRoleResponse(role *model.Role) *response.RoleResponse {
	if role == nil {
		return nil
	}
	return &response.RoleResponse{
		ID:      role.ID,
		Name:    role.Name,
		Code:    role.Code,
		Status:  role.Status,
		Sort:    role.Sort,
		Remark:  role.Remark,
		Created: role.CreatedAt.Format(time.DateTime),
		Updated: role.UpdatedAt.Format(time.DateTime),
	}
}

func ToRoleList(roles []*model.Role) []*response.RoleResponse {
	var list []*response.RoleResponse
	for _, role := range roles {
		list = append(list, ToRoleResponse(role))
	}
	return list
}
