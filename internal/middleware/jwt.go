package middleware

import (
	"github.com/wxlbd/nunu-casbin-admin/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth(jwt *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "未登录或非法访问",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "请求头Authorization格式错误",
			})
			c.Abort()
			return
		}

		token = token[7:]
		claims, err := jwt.ParseToken(c, token, false)
		if err != nil {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "token已过期或非法token",
			})
			c.Abort()
			return
		}

		// 检查是否需要续期
		newToken, needRenew, err := jwt.CheckAndRenewToken(c, token, claims)
		if err != nil {
			c.JSON(500, gin.H{
				"code":    500,
				"message": "token续期失败",
			})
			c.Abort()
			return
		}

		// 如果需要续期，在响应头中返回新的token
		if needRenew {
			c.Header("New-Access-Token", newToken)
		}

		// 将用户信息存入上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
