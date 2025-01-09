package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前登录用户
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "未登录",
			})
			c.Abort()
			return
		}

		// 获取请求的URI和方法
		obj := c.Request.URL.Path
		act := c.Request.Method

		// 检查权限
		ok, err := enforcer.Enforce(user, obj, act)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "权限检查失败",
			})
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": "没有权限",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
