package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
)

// ErrorHandler 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否有错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var e *errors.Error

			// 处理验证错误
			if validationError := errors.ParseValidateError(err); validationError != nil {
				e = validationError
			} else if customErr, ok := err.(*errors.Error); ok {
				// 处理自定义错误
				e = customErr
			} else {
				// 处理其他错误
				e = errors.ErrUnknown.WithMessage(err.Error())
			}

			// 返回错误响应
			c.JSON(e.Status, gin.H{
				"code":    e.Code,
				"message": e.Message,
			})
			c.Abort()
		}
	}
}
