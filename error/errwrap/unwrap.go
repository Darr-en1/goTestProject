package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// Unwrap 解除封装并进行断言处理
func Unwrap() {

	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Println("wrapped error: ", err)

	// 处理错误类型
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred: ", err)
	default:
		fmt.Println("an unknown error occurred")
	}
}

// StackTrace 打印错误栈
func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}

//func RetError() error{
//	err := ThisReturnsAnError()
//	return errors.Wrap(err, "This only does something if err != nil")
//}
