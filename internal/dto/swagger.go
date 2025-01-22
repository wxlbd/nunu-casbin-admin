package dto

// 分页响应
type PageResponse struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}

// 通用响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
