package parser

import (
	"goTestProject/crewler/engine"
	"goTestProject/crewler/model"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	stringSubmatch := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for idx, strings := range stringSubmatch {
		result.Items = append(result.Items, model.City{Name: string(strings[2])})
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(strings[1]),
				ParseFunc: ParseCity,
			})
		if idx == 0 {
			break
		}
	}
	return result
}
