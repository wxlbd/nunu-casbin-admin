package model

import "time"

type DictData struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	Sort      int    `json:"sort" gorm:"column:dict_sort"`
	Label     string `json:"label" gorm:"column:label"`
	Value     string `json:"value" gorm:"column:value"`
	TypeCode  string `json:"type_code" gorm:"column:type_code"`
	Status    int    `json:"status" gorm:"column:status"`
	Remark    string `json:"remark" gorm:"column:remark"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
