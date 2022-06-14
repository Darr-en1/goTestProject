package error_handle

import (
	"github.com/pkg/errors"
)

type ErrCode int64

// CustomError 自定义error结构体，并重写Error()方法
// 错误时返回自定义结构
type CustomError struct {
	Code    ErrCode `json:"code"`    // 业务码
	Message string  `json:"message"` // 业务码
}

func (e *CustomError) Error() string {
	return e.Code.String()
}

// 定义errorCode
//go:generate stringer -type ErrCode -linecomment
const (
	// 服务级错误码
	ServerError        ErrCode = 10101 // Internal Server Error
	TooManyRequests    ErrCode = 10102 // Too Many Requests
	ParamBindError     ErrCode = 10103 // 参数信息有误
	AuthorizationError ErrCode = 10104 // 签名信息有误
	CallHTTPError      ErrCode = 10105 // 调用第三方HTTP接口失败
	ResubmitError      ErrCode = 10106 // ResubmitError
	ResubmitMsg        ErrCode = 10107 // 请勿重复提交
	HashIdsDecodeError ErrCode = 10108 // ID参数有误
	SignatureError     ErrCode = 10109 // SignatureError
)

// NewCustomError 新建自定义error实例化
func NewCustomError(code ErrCode) error {
	// 初次调用得用Wrap方法，进行实例化
	return errors.Wrap(&CustomError{
		Code:    code,
		Message: code.String(),
	}, "")
}
