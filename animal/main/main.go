package main

import (
	"fmt"
	"time"
)

var i = 60

func main() {
	fmt.Println(i * time.Second)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//只有最后一次panic调用被捕获
	defer func() {
		panic("first defer panic") //打印结构是这个
	}()
	defer func() {
		panic("second defer panic")
	}()
	panic("main body panic")
}
