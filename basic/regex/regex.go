package main

import (
	"fmt"
	"regexp"
)

const text = `My email is darr_en1@126.com
				email1 is darr_en1@126.com.cn`

func main() {
	// ``不会被转义 ，可以把 () 去掉，()用于 Submatch
	re := regexp.MustCompile(`([a-zA-Z0-9_]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	match := re.FindString(text)

	matchSlice := re.FindAllString(text, -1)

	// 提取：匹配()的内容
	subMatch := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match, matchSlice, subMatch)

}
