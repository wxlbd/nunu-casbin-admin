package ginx

import "github.com/gin-gonic/gin"

const (
	SUCCESS = 200
	ERROR   = 500
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data any) {
	c.JSON(200, Response{
		Code:    SUCCESS,
		Message: "success",
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

// ParamError 参数错误响应
func ParamError(c *gin.Context) {
	Error(c, 400, "参数错误")
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context) {
	Error(c, 401, "未授权")
}

// Forbidden 禁止访问响应
func Forbidden(c *gin.Context) {
	Error(c, 403, "禁止访问")
}

// ServerError 服务器错误响应
func ServerError(c *gin.Context, err error) {
	Error(c, 500, err.Error())
}
