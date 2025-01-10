package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Menu 菜单模型
type Menu struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	ParentID  uint64    `json:"parent_id" gorm:"default:0"`  // 父菜单ID，0表示根菜单
	Name      string    `json:"name" gorm:"size:64"`         // 菜单名称
	Path      string    `json:"path" gorm:"size:128"`        // 路由路径
	Component string    `json:"component" gorm:"size:128"`   // 组件路径
	Sort      int16     `json:"sort" gorm:"default:0"`       // 排序，值越小越靠前
	Status    int8      `json:"status" gorm:"default:1"`     // 1: 正常, 2: 禁用
	Redirect  string    `json:"redirect" gorm:"size:128"`    // 重定向路径
	Meta      MenuMeta  `json:"meta" gorm:"type:json"`       // 菜单元数据
	CreatedBy uint64    `json:"created_by" gorm:"default:0"` // 创建者
	UpdatedBy uint64    `json:"updated_by" gorm:"default:0"` // 更新者
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Remark    string    `json:"remark" gorm:"size:255"` // 备注
}

// MenuMeta 菜单元数据
type MenuMeta struct {
	I18n             string `json:"i18n"`             // 国际化标识
	Icon             string `json:"icon"`             // 图标
	Type             string `json:"type"`             // 类型：M=菜单,B=按钮（按钮类型对应的就是后端API）
	Affix            bool   `json:"affix"`            // 是否固定标签
	Cache            bool   `json:"cache"`            // 是否缓存
	Title            string `json:"title"`            // 标题
	Hidden           bool   `json:"hidden"`           // 是否隐藏
	Copyright        bool   `json:"copyright"`        // 是否有版权
	ComponentPath    string `json:"componentPath"`    // 组件路径
	ComponentSuffix  string `json:"componentSuffix"`  // 组件后缀
	BreadcrumbEnable bool   `json:"breadcrumbEnable"` // 是否启用面包屑
}

// TableName 指定表名
func (Menu) TableName() string {
	return "menu"
}

// Scan 实现 sql.Scanner 接口，用于从数据库读取 JSON 数据到结构体
func (m *MenuMeta) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSON value: value is not a byte slice")
	}

	return json.Unmarshal(bytes, m)
}

// Value 实现 driver.Valuer 接口，用于将结构体序列化为 JSON 存储到数据库
func (m MenuMeta) Value() (driver.Value, error) {
	return json.Marshal(m)
}
