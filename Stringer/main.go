package main

import (
	"fmt"
	"goTestProject/Stringer/error_handle"
)

func main() {
	fmt.Println(error_handle.NewCustomError(error_handle.AuthorizationError))
	fmt.Println(error_handle.AuthorizationError)
}
