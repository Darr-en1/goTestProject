package basicerrors

import (
	"errors"
	"fmt"
)

// ErrorValue ErrorValue创建了包级错误
// 可以采用这样的方式判断： if err == ErrorValue
var ErrorValue = errors.New("this is a typed error")

// TypedError TypedError创建了包含错误类型的匿名字段
// 可以采用断言的方式判断：err.(type) == ErrorValue
type TypedError struct {
	error
}

//BasicErrors 演示了错误的创建
func BasicErrors() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occurred: %s", "something")
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorValue
	fmt.Println("value error: ", err)

	err = TypedError{errors.New("typed error")}
	fmt.Println("typed error: ", err)

}
