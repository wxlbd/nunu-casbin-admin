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
	Page     int `json:"page" form:"page" binding:"required,min=1"`
	PageSize int `json:"pageSize" form:"page_size" binding:"required,min=1,max=100"`
}
