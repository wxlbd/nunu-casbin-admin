package model

// UserRoles 用户-角色关联模型
type UserRoles struct {
	ID     uint64 `gorm:"primaryKey"`
	UserID uint64
	RoleID uint64
}

// TableName 指定表名
func (UserRoles) TableName() string {
	return "user_roles"
}
