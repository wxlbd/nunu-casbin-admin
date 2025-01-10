package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wxlbd/nunu-casbin-admin/pkg/log"
	"go.uber.org/zap"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// RequestLogger 请求日志中间件
func RequestLogger(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 记录请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 包装 ResponseWriter
		w := &responseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = w

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()
		latency := end.Sub(start)

		// 获取用户信息
		userID, _ := c.Get("user_id")
		username, _ := c.Get("username")

		// 构建日志字段
		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
		}

		// 添加用户信息
		if userID != nil {
			fields = append(fields, zap.Any("user_id", userID))
		}
		if username != nil {
			fields = append(fields, zap.Any("username", username))
		}

		// 添加请求体（如果不是GET请求）
		if c.Request.Method != "GET" && len(requestBody) > 0 {
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, requestBody, "", "  "); err == nil {
				fields = append(fields, zap.String("request", prettyJSON.String()))
			} else {
				fields = append(fields, zap.ByteString("request", requestBody))
			}
		}

		// 添加响应体
		if w.body.Len() > 0 {
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, w.body.Bytes(), "", "  "); err == nil {
				fields = append(fields, zap.String("response", prettyJSON.String()))
			} else {
				fields = append(fields, zap.String("response", w.body.String()))
			}
		}

		// 记录日志
		if c.Writer.Status() >= 500 {
			logger.Error("Request failed", fields...)
		} else {
			logger.Info("Request completed", fields...)
		}
	}
}
