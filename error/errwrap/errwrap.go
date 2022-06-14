package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// WrappedError 演示了如何对错误进行封装
func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

type ErrorTyped struct {
	error
}

func Wrap() {
	e := errors.New("standard error")

	fmt.Println("Regular Error - ", WrappedError(e))
	//fmt.Printf("Regular Error - %+v \n", WrappedError(e))

	fmt.Println("Typed Error - ", WrappedError(ErrorTyped{errors.New("typed error")}))

	fmt.Println("Nil -", WrappedError(nil))

}
