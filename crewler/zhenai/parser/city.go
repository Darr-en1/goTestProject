package parser

import (
	"goTestProject/crewler/engine"
	"goTestProject/crewler/model"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	personSubmatch := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for idx, person := range personSubmatch {
		profile := model.Profile{Name: string(person[2])}
		result.Items = append(result.Items, profile)
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(person[1]),
				ParseFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(bytes, profile)
				},
			})
		if idx == 0 {
			//break
		}
	}
	return result
}
