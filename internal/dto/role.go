package dto

import (
	"github.com/wxlbd/nunu-casbin-admin/internal/model"
	"time"
)

func ToRoleResponse(role *model.Role) *RoleResponse {
	if role == nil {
		return nil
	}
	return &RoleResponse{
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

func ToRoleList(roles []*model.Role) []*RoleResponse {
	var list []*RoleResponse
	for _, role := range roles {
		list = append(list, ToRoleResponse(role))
	}
	return list
}

// RoleRequest 创建/更新角色请求
type RoleRequest struct {
	ID     uint64 `json:"id"` // 更新时必填
	Name   string `json:"name" binding:"required"`
	Code   string `json:"code" binding:"required"`
	Status int8   `json:"status"`
	Sort   int16  `json:"sort"`
	Remark string `json:"remark"`
}

// RoleIDRequest 删除角色请求
type RoleIDRequest struct {
	ID uint64 `json:"id" binding:"required"`
}

// RoleListRequest 角色列表请求
type RoleListRequest struct {
	Page int `form:"page" binding:"required,min=1"`
	Size int `form:"size" binding:"required,min=1,max=100"`
}

// AssignMenusRequest 分配菜单请求
type AssignMenusRequest struct {
	RoleID  uint64   `json:"role_id" binding:"required"`
	MenuIDs []uint64 `json:"menu_ids" binding:"required"`
}

// RoleResponse 角色信息响应
type RoleResponse struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Status  int8   `json:"status"`
	Sort    int16  `json:"sort"`
	Remark  string `json:"remark"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

// RoleListResponse 角色列表响应
type RoleListResponse struct {
	List  []*RoleResponse `json:"list"`
	Total int64           `json:"total"`
}

// RoleMenusResponse 角色菜单响应
type RoleMenusResponse []*Menu
