package engine

import (
	"goTestProject/crewler/fetcher"
	"log"
)

func Run(seeds ...Request) { //... :https://blog.csdn.net/jeffrey11223/article/details/79166724
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %+v", item)
		}

	}
}
