package recommender

import (
	"strconv"

	"nimaajdari.com/nobadweather/api"
	"nimaajdari.com/nobadweather/errs"
)

type GeoApi interface {
	SearchLocation(query string) ([]api.SearchResult, error)
}

type WeatherApi interface {
	GetWeatherInfo(lat, lon float64) ([]api.WeatherData, error)
}

type LineReader interface {
	ReadLine() (line []byte, isPrefix bool, err error)
}

type Cli struct {
	geoApi     GeoApi
	weatherApi WeatherApi
	reader     LineReader
	println    func(a ...any) (n int, err error)
	printf     func(format string, a ...any) (n int, err error)
}

func NewCli(geoApi GeoApi,
	weatherApi WeatherApi,
	reader LineReader,
	println func(a ...any) (n int, err error),
	printf func(format string, a ...any) (n int, err error)) *Cli {
	return &Cli{geoApi: geoApi, weatherApi: weatherApi, reader: reader, println: println, printf: printf}
}

func (cli Cli) Run() error {
	var address []byte
	var geoSearchResult []api.SearchResult
	cli.println("Enter your location address or city:")

	for len(address) == 0 || len(geoSearchResult) == 0 {
		var err error
		address, _, err = cli.reader.ReadLine()
		if err != nil {
			return errs.StackInfo(err)
		}
		if len(address) == 0 {
			continue
		}
		geoSearchResult, err = cli.geoApi.SearchLocation(string(address))
		if err != nil {
			return errs.StackInfo(err)
		}
		if len(geoSearchResult) == 0 {
			cli.println("No location found for your search query. Try again.")
		}
	}

	cli.println("Select the address by entering the number:")
	for i, v := range geoSearchResult {
		cli.printf("%d: %s \n", i+1, v)
	}

	var indexSelectionInput []byte
	var selectionNumber int
	for len(indexSelectionInput) == 0 || selectionNumber == 0 {
		var err error

		indexSelectionInput, _, err = cli.reader.ReadLine()
		if err != nil {
			return errs.StackInfo(err)
		}
		if len(indexSelectionInput) == 0 {
			continue
		}
		selectionNumber, err = strconv.Atoi(string(indexSelectionInput))
		if err != nil {
			cli.println("Failed to parse input as int.")
			selectionNumber = 0
		}
		if selectionNumber < 1 || selectionNumber > len(geoSearchResult) {
			cli.println("Invalid index selection.")
			selectionNumber = 0
		}
	}
	locSelection := geoSearchResult[selectionNumber-1]

	weatherData, err := cli.weatherApi.GetWeatherInfo(locSelection.Lat, locSelection.Lon)
	if err != nil {
		return errs.StackInfo(err)
	}

	recommendation, err := RecommendClothing(weatherData)
	if err != nil {
		return errs.StackInfo(err)
	}

	cli.println(recommendation)
	return nil
}
