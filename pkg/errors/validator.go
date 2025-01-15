package errors

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateErrorMsg 定义验证错误信息映射
var ValidateErrorMsg = map[string]string{
	"required": "不能为空",
	"email":    "必须是有效的电子邮件地址",
	"min":      "不能小于 %s",
	"max":      "不能大于 %s",
	"len":      "长度必须是 %s",
	"oneof":    "必须是 [%s] 中的一个",
}

// ParseValidateError 解析验证错误
func ParseValidateError(err error) *Error {
	if err == nil {
		return nil
	}

	// 处理验证器错误
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		// 如果只有一个错误，直接返回
		if len(validationErrors) == 1 {
			return New(ValidationFailed, formatValidateError(validationErrors[0]))
		}

		// 多个错误，组合成一个消息
		var errMsgs []string
		for _, e := range validationErrors {
			errMsgs = append(errMsgs, formatValidateError(e))
		}
		return New(ValidationFailed, strings.Join(errMsgs, "; "))
	}

	// 处理 JSON 解析错误
	return New(InvalidParam, "请求参数格式错误")
}

// formatValidateError 格式化单个验证错误
func formatValidateError(err validator.FieldError) string {
	field := err.Field()
	tag := err.Tag()
	param := err.Param()

	// 获取错误信息模板
	template, exists := ValidateErrorMsg[tag]
	if !exists {
		template = "验证失败"
	}

	// 如果有参数，格式化错误信息
	if param != "" {
		return fmt.Sprintf("%s %s", field, fmt.Sprintf(template, param))
	}

	return fmt.Sprintf("%s %s", field, template)
}
