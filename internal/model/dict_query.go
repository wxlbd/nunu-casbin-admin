package model

type DictTypeQuery struct {
	Name     string `form:"name"`
	Code     string `form:"code"`
	Status   int32  `form:"status"`
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"page_size,default=10"`
}

type DictDataQuery struct {
	TypeCode string `form:"type_code"`
	Label    string `form:"label"`
	Status   int32  `form:"status"`
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"page_size,default=10"`
}