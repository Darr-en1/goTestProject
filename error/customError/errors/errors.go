package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type ErrorType uint

//go:generate stringer -type ErrorType -linecomment
const (
	NoType     = ErrorType(iota) // 未知异常
	BadRequest                   // 请求错误
	NotFound                     //  没有资源

	//增加任何你想要的类型
)

type customError struct {
	errorType     ErrorType
	originalError error
	contextInfo   map[string]string
}

// Error 方法返回一个 customError 消息
func (error customError) Error() string {
	return error.originalError.Error()
}

// New 方法新建一个新的 customError 对象 携带 stack
func (e ErrorType) New(msg string) error {
	return customError{errorType: e, originalError: errors.New(msg)}
}

// Newf 方法使用格式化消息新建 customError 对象
func (e ErrorType) Newf(msg string, args ...interface{}) error {
	newErr := fmt.Errorf(msg, args...)
	return customError{errorType: e, originalError: newErr}
}

// Wrap 方法新建一个封装错误  携带 stack
func (e ErrorType) Wrap(err error, msg string) error {
	return e.Wrapf(err, msg)
}

// Wrapf 方法使用格式化消息创建新的封装错误 携带 stack
func (e ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args...)

	return customError{errorType: e, originalError: newErr}
}

// New 方法新建一个错误类型
func New(msg string) error {
	return customError{errorType: NoType, originalError: errors.New(msg)}
}

// Newf 方法用格式化消息新建了一个错误类型
func Newf(msg string, args ...interface{}) error {
	return customError{errorType: NoType, originalError: errors.Errorf(msg, args...)}
}

// Wrap 方法用字符串封装错误
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Cause 方法返回原始错误
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf 方法用格式化字符串封装错误
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	var customErr customError
	if errors.As(err, &customErr) {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			contextInfo:   customErr.contextInfo,
		}
	}

	return customError{errorType: NoType, originalError: wrappedError}
}

// AddErrorContext 方法为错误添加上下文
func AddErrorContext(err error, field, message string) error {
	var customErr customError
	if errors.As(err, &customErr) {
		customErr.contextInfo[field] = message
		return customError{errorType: customErr.errorType, originalError: customErr.originalError, contextInfo: customErr.contextInfo}
	}
	return customError{errorType: NoType, originalError: err, contextInfo: map[string]string{field: message}}
}

// GetErrorContext 方法返回错误内容
func GetErrorContext(err error) map[string]string {
	var customErr customError
	if errors.As(err, &customErr) {
		return customErr.contextInfo
	}
	return nil
}

// GetType 方法返回错误类型
func GetType(err error) ErrorType {
	var customErr customError
	if errors.As(err, &customErr) {
		return customErr.errorType
	}
	return NoType
}
