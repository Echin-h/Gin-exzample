package errs

import (
	"errors"
	"fmt"
)

// 下面展现如何生成“错误”的过程

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Origin  string `json:"origin"` // Origin字段通常用于只是错误的来源或者上下文
}

// 绑定数字和文字之间的关联
func newError(code int64, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

// 表示数字背后的含义（人能看懂）
func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Is(target error) bool {
	var t *Error
	// errors.As是把target转化成t的类型
	ok := errors.As(target, &t)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

// 添加error类型
func (e *Error) WithOrigin(err error) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message,
		Origin:  fmt.Sprintf("%+v", err),
	}
}

// 添加字符串进入
func (e *Error) WithTips(details ...string) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message + " " + fmt.Sprintf("%v", details),
	}
}
