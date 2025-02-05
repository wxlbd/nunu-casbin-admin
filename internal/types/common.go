package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
)

// BackendSetting 后台设置结构体
type BackendSetting struct {
	App struct {
		Layout          string   `json:"layout"`
		AsideDark       bool     `json:"asideDark"`
		ColorMode       string   `json:"colorMode"`
		UseLocale       string   `json:"useLocale"`
		WhiteRoute      []string `json:"whiteRoute"`
		PageAnimate     string   `json:"pageAnimate"`
		PrimaryColor    string   `json:"primaryColor"`
		WatermarkText   string   `json:"watermarkText"`
		ShowBreadcrumb  bool     `json:"showBreadcrumb"`
		EnableWatermark bool     `json:"enableWatermark"`
		LoadUserSetting bool     `json:"loadUserSetting"`
	} `json:"app,omitempty"`
	Tabbar struct {
		Mode   string `json:"mode"`
		Enable bool   `json:"enable"`
	} `json:"tabbar,omitempty"`
	SubAside struct {
		ShowIcon           bool `json:"showIcon"`
		ShowTitle          bool `json:"showTitle"`
		FixedAsideState    bool `json:"fixedAsideState"`
		ShowCollapseButton bool `json:"showCollapseButton"`
	} `json:"subAside"`
	Copyright struct {
		Dates       string `json:"dates"`
		Enable      bool   `json:"enable"`
		Company     string `json:"company"`
		Website     string `json:"website"`
		PutOnRecord string `json:"putOnRecord"`
	} `json:"copyright"`
	MainAside struct {
		ShowIcon             bool `json:"showIcon"`
		ShowTitle            bool `json:"showTitle"`
		EnableOpenFirstRoute bool `json:"enableOpenFirstRoute"`
	} `json:"mainAside"`
	WelcomePage struct {
		Icon  string `json:"icon"`
		Name  string `json:"name"`
		Path  string `json:"path"`
		Title string `json:"title"`
	} `json:"welcomePage"`
}

// Value 实现 driver.Valuer 接口
func (b BackendSetting) Value() (driver.Value, error) {
	return json.Marshal(b)
}

// Scan 实现 sql.Scanner 接口
func (b *BackendSetting) Scan(src any) error {
	if src == nil {
		*b = BackendSetting{} // 清空当前结构体
		return nil
	}

	jsonData, ok := src.([]byte)
	if !ok {
		return errors.New("invalid data type for BackendSetting")
	}

	// 检查是否为空字符串、空对象或空数组
	if len(jsonData) == 0 || string(jsonData) == "{}" || string(jsonData) == "[]" {
		*b = BackendSetting{} // 清空当前结构体
		return nil
	}

	// 使用别名类型避免递归调用
	type Alias BackendSetting
	aux := (*Alias)(b)
	return json.Unmarshal(jsonData, aux)
}

// MarshalJSON 自定义 JSON 序列化
func (b *BackendSetting) MarshalJSON() ([]byte, error) {
	if b == nil || reflect.DeepEqual(*b, BackendSetting{}) {
		return []byte("null"), nil
	}

	// 使用别名类型避免递归调用
	type Alias BackendSetting
	return json.Marshal((*Alias)(b))
}

// UnmarshalJSON 自定义 JSON 反序列化
func (b *BackendSetting) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" || string(data) == "{}" || string(data) == "[]" {
		*b = BackendSetting{} // 清空当前结构体
		return nil
	}

	// 使用别名类型避免递归调用
	type Alias BackendSetting
	aux := (*Alias)(b)
	return json.Unmarshal(data, aux)
}

// PageParam 分页请求参数
type PageParam struct {
	Page     int `json:"page" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Normalize 规范化分页参数
func (p *PageParam) Normalize() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
}

// GetOffset 获取偏移量
func (p *PageParam) GetOffset() int {
	return (p.Page - 1) * p.PageSize
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
