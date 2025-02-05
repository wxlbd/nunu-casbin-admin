package errors

import "net/http"

var (
	// 通用错误
	ErrUnknown      = New(Unknown, "未知错误").WithStatus(http.StatusInternalServerError)
	ErrInvalidParam = New(InvalidParam, "参数错误").WithStatus(http.StatusBadRequest)
	ErrUnauthorized = New(Unauthorized, "未授权").WithStatus(http.StatusUnauthorized)
	ErrForbidden    = New(Forbidden, "禁止访问").WithStatus(http.StatusForbidden)
	ErrNotFound     = New(NotFound, "资源不存在").WithStatus(http.StatusNotFound)
	ErrDatabase     = New(DatabaseError, "数据库错误").WithStatus(http.StatusInternalServerError)
	// ErrTokenExpired token过期
	ErrTokenExpired = New(TokenExpired, "令牌过期").WithStatus(http.StatusUnauthorized)
	// ErrTokenInvalid token无效
	ErrTokenInvalid = New(TokenInvalid, "令牌无效").WithStatus(http.StatusUnauthorized)

	// 业务错误
	ErrCircularReference = New(CircularReference, "检测到循环引用")
	ErrDuplicateEntry    = New(DuplicateEntry, "数据已存在")
	ErrStatusConflict    = New(StatusConflict, "状态冲突")
)

// WithMsg 创建新的错误消息
func WithMsg(code ErrorCode, msg string) *Error {
	return New(code, msg)
}
