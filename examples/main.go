package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/solar-jsoc/gohunter"
)

func main() {
	hunterClient := gohunter.NewClient("*your token*", gohunter.WithCustomClient(&http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
		Timeout: 5 * time.Minute,
	}))

	result, err := hunterClient.DomainSearch(context.Background(), "*your domain*", "",
		gohunter.WithLimit(20),
		gohunter.WithSeniority(gohunter.SenioritySenior))
	if err != nil {
		if e, ok := err.(*gohunter.HunterError); ok {
			switch e.Code {
			case 429:
				// handle error
			}
		}
	}

	fmt.Printf("%+v\n", result)

}
