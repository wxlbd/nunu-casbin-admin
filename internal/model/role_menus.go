package model

// RoleMenus 角色-菜单关联模型
type RoleMenus struct {
	ID     uint64 `gorm:"primaryKey"`
	RoleID uint64
	MenuID uint64
}

func (RoleMenus) TableName() string {
	return "role_menus"
}
