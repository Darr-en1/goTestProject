package engine

import "fmt"

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{} // interface{} 表示任何类型
}

func NilParser(contents []byte) ParseResult { // 定义一个空函数保证执行不出错
	fmt.Printf("%s", contents)
	return ParseResult{}
}
