package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	//fmt.Println(error_handle2.NewCustomError(error_handle2.AuthorizationError))
	//fmt.Println(error_handle2.AuthorizationError)
	err := errors.Errorf("你好呀", "")
	fmt.Printf("%+v", err)
}
