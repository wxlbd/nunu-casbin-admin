package errors

// ErrorCode 错误码类型
type ErrorCode int

const (
	// 成功
	Success ErrorCode = 0

	// 通用错误 (10000-19999)
	Unknown       ErrorCode = 10000 // 未知错误
	InvalidParam  ErrorCode = 10001 // 参数错误
	Unauthorized  ErrorCode = 10002 // 未授权
	Forbidden     ErrorCode = 10003 // 禁止访问
	NotFound      ErrorCode = 10004 // 资源不存在
	AlreadyExists ErrorCode = 10005 // 资源已存在
	ServerError   ErrorCode = 10006 // 服务器错误
	DatabaseError ErrorCode = 10007 // 数据库错误

	// 业务错误 (20000-99999)
	// 按错误类型分类，而不是按模块分类

	// 认证相关 (20000-20999)
	InvalidCredentials ErrorCode = 20000 // 凭证无效
	TokenExpired       ErrorCode = 20001 // 令牌过期
	TokenInvalid       ErrorCode = 20002 // 令牌无效

	// 权限相关 (21000-21999)
	PermissionDenied ErrorCode = 21000 // 权限不足
	RoleNotAssigned  ErrorCode = 21001 // 未分配角色

	// 数据验证相关 (22000-22999)
	ValidationFailed ErrorCode = 22000 // 验证失败
	DuplicateEntry   ErrorCode = 22001 // 重复数据
	InvalidFormat    ErrorCode = 22002 // 格式错误

	// 业务规则相关 (23000-23999)
	BusinessRuleViolation ErrorCode = 23000 // 违反业务规则
	CircularReference     ErrorCode = 23001 // 循环引用
	StatusConflict        ErrorCode = 23002 // 状态冲突

	// 外部服务相关 (24000-24999)
	ExternalServiceError ErrorCode = 24000 // 外部服务错误
	ServiceTimeout       ErrorCode = 24001 // 服务超时
)
