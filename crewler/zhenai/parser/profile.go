package parser

import (
	"goTestProject/crewler/engine"
	"goTestProject/crewler/model"
	"regexp"
)

const profileRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func ParseProfile(contents []byte, profile model.Profile) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	profileSubmatch := re.FindSubmatch(contents)
	if profileSubmatch != nil {
		profile.Home = string(profileSubmatch[1])
	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
