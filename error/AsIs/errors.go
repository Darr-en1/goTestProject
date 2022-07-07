package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

type NewError struct {
	err error
	msg string
}

func (e *NewError) Error() string {
	return e.msg + e.err.Error()
}

func main() {
	_, err := div(10, 0)
	newErr := fmt.Errorf("div 错误：%s error", err)
	// %s 没有 Wrapping Error
	fmt.Println(err == ErrDivByZero)
	fmt.Println(newErr == ErrDivByZero)

	// 通常只能文本比较
	fmt.Println(strings.Contains(err.Error(), ErrDivByZero.Error()))
	fmt.Println(strings.Contains(newErr.Error(), ErrDivByZero.Error()))

	// 封装error 从而保留原有error
	newErr1 := NewError{err, "执行出错："}
	fmt.Println(newErr1.Error())
	fmt.Println(newErr1.err == ErrDivByZero)

	warpErr := fmt.Errorf("div 错误：%w error", err)
	// %w  实现 Wrapping Error,通过 errors.Unwrap(w) 获取 原始error
	fmt.Println(err == ErrDivByZero)
	fmt.Println(warpErr == ErrDivByZero)
	fmt.Println(errors.Unwrap(warpErr) == ErrDivByZero)

	// Is 判断Wrapping Error 原始error
	//fmt.Println(errorss(warpErr, ErrDivByZero))

	//err1 := errorss.Wrap(err, "11")
	//fmt.Println(err == ErrDivByZero)
	//fmt.Println(err1 == ErrDivByZero)
	//fmt.Println(errors.Is(err1, ErrDivByZero))
	fmt.Println(errors.As(warpErr, &ErrDivByZero))

	//switch z, err := div(10, 0); err {
	//case nil:
	//	println(z)
	//case ErrDivByZero:
	//	panic(err)
	//}
}
