package main

import (
	"errors"
	"fmt"
	errorss "github.com/pkg/errors"
)

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}
func main() {
	_, err := div(10, 0)
	newErr := fmt.Errorf("div 错误：%s error", err)
	fmt.Println(err == ErrDivByZero)
	fmt.Println(newErr == ErrDivByZero)

	//fmt.Println(aa == ErrDivByZero)
	//fmt.Println(errorss.Is(a,ErrDivByZero))
	//fmt.Println(errorss.Is(aa,ErrDivByZero))

	err1 := errorss.Wrap(err, "11")
	fmt.Println(err == ErrDivByZero)
	fmt.Println(err1 == ErrDivByZero)
	fmt.Println(errors.Is(err1, ErrDivByZero))
	fmt.Println(errors.As(err1, &ErrDivByZero))

	//switch z, err := div(10, 0); err {
	//case nil:
	//	println(z)
	//case ErrDivByZero:
	//	panic(err)
	//}
}
