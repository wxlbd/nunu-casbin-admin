package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// User 用户模型
type User struct {
	ID             uint64         `json:"id" gorm:"primaryKey"`
	Username       string         `json:"username" gorm:"uniqueIndex;size:64"`
	Password       string         `json:"-" gorm:"size:128"`
	Nickname       string         `json:"nickname" gorm:"size:128"`
	Phone          string         `json:"phone" gorm:"size:16"`
	Email          string         `json:"email" gorm:"size:128"`
	Avatar         string         `json:"avatar" gorm:"size:255"`
	Status         int8           `json:"status" gorm:"default:1"`
	UserType       int            `json:"user_type" gorm:"default:0"`
	Signed         string         `json:"signed" gorm:"size:255"`
	LoginIp        string         `json:"login_ip" gorm:"size:64"`
	LoginTime      time.Time      `json:"login_time"`
	BackendSetting BackendSetting `json:"backend_setting" gorm:"type:json"`
	CreatedBy      uint64         `json:"created_by" gorm:"default:0"`
	UpdatedBy      uint64         `json:"updated_by" gorm:"default:0"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	Remark         string         `json:"remark" gorm:"size:255"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}

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
	} `json:"app"`
	Tabbar struct {
		Mode   string `json:"mode"`
		Enable bool   `json:"enable"`
	} `json:"tabbar"`
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

func (b BackendSetting) Value() (driver.Value, error) {
	// 将 BackendSetting 结构体转换为 JSON 格式的字节切片
	jsonData, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (b *BackendSetting) Scan(src any) error {
	// 检查 src 是否为 nil
	if src == nil {
		return nil
	}

	// 将 src 转换为字节切片
	jsonData, ok := src.([]byte)
	if !ok {
		return errors.New("invalid data type for BackendSetting")
	}

	// 将 JSON 格式的字节切片解析为 BackendSetting 结构体
	return json.Unmarshal(jsonData, b)
}
