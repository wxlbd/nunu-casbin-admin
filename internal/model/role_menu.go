package model

// RoleBelongsMenu 角色-菜单关联模型
type RoleBelongsMenu struct {
	RoleID uint64 `gorm:"primaryKey"`
	MenuID uint64 `gorm:"primaryKey"`
}

// TableName 指定表名
func (RoleBelongsMenu) TableName() string {
	return "role_belongs_menu"
}
