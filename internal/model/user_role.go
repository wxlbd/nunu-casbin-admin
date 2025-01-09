package model

// UserBelongsRole 用户-角色关联模型
type UserBelongsRole struct {
	UserID uint64 `gorm:"primaryKey"`
	RoleID uint64 `gorm:"primaryKey"`
}

// TableName 指定表名
func (UserBelongsRole) TableName() string {
	return "user_belongs_role"
}
