package main

import (
	"goTestProject/crewler/engine"
	"goTestProject/crewler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
