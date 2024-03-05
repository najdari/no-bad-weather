package main

import (
	"bufio"
	"fmt"
	"os"

	"nimaajdari.com/nobadweather/api"
	"nimaajdari.com/nobadweather/recommender"
)

func main() {
	geoApiKey := os.Getenv("GEOAPIFY_API_KEY")
	if geoApiKey == "" {
		panic("no GEOAPIFY_API_KEY set")
	}

	yrUserAgent := os.Getenv("YR_API_USER_AGENT")
	if yrUserAgent == "" {
		panic("no YR_API_USER_AGENT set")
	}

	reader := bufio.NewReader(os.Stdin)

	cli := recommender.NewCli(api.NewGeoapifyClient(geoApiKey), api.NewWeatherApi(yrUserAgent), reader, fmt.Println, fmt.Printf)
	err := cli.Run()
	if err != nil {
		panic(err)
	}
}
