package dto

import (
	"time"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
)

// DictType DTOs
type DictTypeRequest struct {
	ID     int64  `json:"id"`
	Name   string `json:"name" binding:"required"`
	Code   string `json:"code" binding:"required"`
	Status int32  `json:"status" binding:"required"`
	Sort   int32  `json:"sort"`
	Remark string `json:"remark"`
}

type DictTypeResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Status    int32     `json:"status"`
	Sort      int32     `json:"sort"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DictData DTOs
type DictDataRequest struct {
	ID       int64  `json:"id"`
	TypeCode string `json:"typeCode" binding:"required"`
	Label    string `json:"label" binding:"required"`
	Value    string `json:"value" binding:"required"`
	Status   int32  `json:"status" binding:"required"`
	Sort     int32  `json:"sort"`
	Remark   string `json:"remark"`
}

type DictDataResponse struct {
	ID        int64     `json:"id"`
	TypeCode  string    `json:"type_code"`
	Label     string    `json:"label"`
	Value     string    `json:"value"`
	Status    int32     `json:"status"`
	Sort      int32     `json:"sort"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 转换函数
func ToDictTypeResponse(dict *model.DictType) *DictTypeResponse {
	return &DictTypeResponse{
		ID:        dict.ID,
		Name:      dict.Name,
		Code:      dict.Code,
		Status:    dict.Status,
		Sort:      dict.Sort,
		Remark:    dict.Remark,
		CreatedAt: dict.CreatedAt,
		UpdatedAt: dict.UpdatedAt,
	}
}

func ToDictDataResponse(data *model.DictDatum) *DictDataResponse {
	return &DictDataResponse{
		ID:        data.ID,
		TypeCode:  data.TypeCode,
		Label:     data.Label,
		Value:     data.Value,
		Status:    data.Status,
		Sort:      data.Sort,
		Remark:    data.Remark,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
