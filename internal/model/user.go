package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex;size:64"`
	Password  string    `json:"-" gorm:"size:128"` // json:"-" 表示不序列化该字段
	Nickname  string    `json:"nickname" gorm:"size:128"`
	Phone     string    `json:"phone" gorm:"size:16"`
	Email     string    `json:"email" gorm:"size:128"`
	Avatar    string    `json:"avatar" gorm:"size:255"`
	Status    int8      `json:"status" gorm:"default:1"` // 1: 正常, 2: 禁用
	LoginTime time.Time `json:"login_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
