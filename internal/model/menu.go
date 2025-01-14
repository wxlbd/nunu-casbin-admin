package model

import (
	"time"

	"github.com/wxlbd/nunu-casbin-admin/internal/types"
)

// Menu 菜单模型
type Menu struct {
	ID        uint64          `json:"id" gorm:"primaryKey"`
	ParentID  uint64          `json:"parent_id" gorm:"default:0"`  // 父菜单ID，0表示根菜单
	Name      string          `json:"name" gorm:"size:64"`         // 菜单名称
	Path      string          `json:"path" gorm:"size:128"`        // 路由路径
	Component string          `json:"component" gorm:"size:128"`   // 组件路径
	Sort      int16           `json:"sort" gorm:"default:0"`       // 排序，值越小越靠前
	Status    int8            `json:"status" gorm:"default:1"`     // 1: 正常, 2: 禁用
	Redirect  string          `json:"redirect" gorm:"size:128"`    // 重定向路径
	Meta      *types.MenuMeta `json:"meta" gorm:"type:json"`       // 菜单元数据
	CreatedBy uint64          `json:"created_by" gorm:"default:0"` // 创建者
	UpdatedBy uint64          `json:"updated_by" gorm:"default:0"` // 更新者
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Remark    string          `json:"remark" gorm:"size:255"` // 备注
}

// TableName 指定表名
func (Menu) TableName() string {
	return "menu"
}

// MenuQuery 菜单查询
type MenuQuery struct {
	types.PageParam
	Name      string `form:"name"`
	Path      string `form:"path"`
	Component string `form:"component"`
	Status    int8   `form:"status"`
	OrderBy   string `form:"order_by"`
}

type MenuTree struct {
	*Menu
	Children []*MenuTree `json:"children"`
}
