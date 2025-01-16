package model

import (
	"time"

	"github.com/wxlbd/gin-casbin-admin/internal/types"
)

// User 用户模型
type User struct {
	ID             uint64                `json:"id" gorm:"primaryKey"`
	Username       string                `json:"username" gorm:"uniqueIndex;size:64"`
	Password       string                `json:"-" gorm:"size:128"`
	Nickname       string                `json:"nickname" gorm:"size:128"`
	Phone          string                `json:"phone" gorm:"size:16"`
	Email          string                `json:"email" gorm:"size:128"`
	Avatar         string                `json:"avatar" gorm:"size:255"`
	Status         int8                  `json:"status" gorm:"default:1"`
	UserType       int                   `json:"user_type" gorm:"default:0"`
	Signed         string                `json:"signed" gorm:"size:255"`
	LoginIp        string                `json:"login_ip" gorm:"size:64"`
	LoginTime      time.Time             `json:"login_time"`
	BackendSetting *types.BackendSetting `json:"backend_setting" gorm:"type:json"`
	CreatedBy      uint64                `json:"created_by" gorm:"default:0"`
	UpdatedBy      uint64                `json:"updated_by" gorm:"default:0"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
	Remark         string                `json:"remark" gorm:"size:255"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}

type UserQuery struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   int8   `json:"status"`
	Nickname string `json:"nickname"`
	Page     int    `json:"page"`
	PageSize int    `json:"size"`
	OrderBy  string `json:"order_by"`
}
