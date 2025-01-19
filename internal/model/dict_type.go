package model

import "time"

type DictType struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Status    int8      `json:"status"`
	Sort      int16     `json:"sort"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
