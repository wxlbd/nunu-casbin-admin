package model

import (
	"time"
)

// Menu 菜单模型
type Menu struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	ParentID  uint64    `json:"parent_id" gorm:"default:0"` // 父菜单ID，0表示根菜单
	Name      string    `json:"name" gorm:"size:64"`        // 菜单名称
	Path      string    `json:"path" gorm:"size:128"`       // 路由路径
	Component string    `json:"component" gorm:"size:128"`  // 组件路径
	Sort      int16     `json:"sort" gorm:"default:0"`      // 排序，值越小越靠前
	Status    int8      `json:"status" gorm:"default:1"`    // 1: 正常, 2: 禁用
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Menu) TableName() string {
	return "menu"
}
