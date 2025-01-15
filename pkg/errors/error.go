package errors

import (
	"fmt"
	"net/http"
)

// Error 自定义错误类型
type Error struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
	Status  int    `json:"-"`       // HTTP状态码
}

func (e *Error) Error() string {
	return e.Message
}

// New 创建新的错误
func New(code ErrorCode, message string) *Error {
	return &Error{
		Code:    int(code),
		Message: message,
		Status:  http.StatusBadRequest, // 默认400
	}
}

// WithStatus 设置HTTP状态码
func (e *Error) WithStatus(status int) *Error {
	e.Status = status
	return e
}

func (e *Error) WithMessage(message string) *Error {
	e.Message = message
	return e
}

// Wrap 包装错误
func Wrap(err error, message string) *Error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok {
		return &Error{
			Code:    e.Code,
			Message: fmt.Sprintf("%s: %s", message, e.Message),
			Status:  e.Status,
		}
	}
	return &Error{
		Code:    int(Unknown),
		Message: fmt.Sprintf("%s: %s", message, err.Error()),
		Status:  http.StatusInternalServerError,
	}
}
