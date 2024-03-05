package recommender_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"nimaajdari.com/nobadweather/api"
	"nimaajdari.com/nobadweather/recommender"
)

type GeoApiMock struct {
	Results []api.SearchResult
}

func (g *GeoApiMock) SearchLocation(query string) ([]api.SearchResult, error) {
	return g.Results, nil
}

type WeatherApiMock struct {
	Results []api.WeatherData
}

func (w *WeatherApiMock) GetWeatherInfo(lat, lon float64) ([]api.WeatherData, error) {
	return w.Results, nil
}

type ReaderMock struct {
	DataSet [][]byte
	index   int
}

func (r *ReaderMock) ReadLine() (line []byte, isPrefix bool, err error) {
	line = r.DataSet[r.index]
	r.index++
	return line, false, nil
}

type StdoutMock struct {
	Out []string
}

func (s *StdoutMock) Println(a ...any) (n int, err error) {
	s.Out = append(s.Out, fmt.Sprint(a...))
	return fmt.Println(a...)
}
func (s *StdoutMock) Printf(format string, a ...any) (n int, err error) {
	s.Out = append(s.Out, fmt.Sprintf(format, a...))
	return fmt.Printf(format, a...)
}

func TestCliWithMockApi(t *testing.T) {

	weatherData := []api.WeatherData{
		{
			AirTemperature:      1.0,
			CloudAreaFraction:   12,
			RelativeHumidity:    13,
			WindSpeed:           4.5,
			PrecipitationAmount: 5,
		},
		{
			AirTemperature:      2.0,
			CloudAreaFraction:   11,
			RelativeHumidity:    15,
			WindSpeed:           5,
			PrecipitationAmount: 2,
		},
	}

	geoApiMock := &GeoApiMock{Results: []api.SearchResult{{Lat: 1, Lon: 1}, {Lat: 2, Lon: 2}}}
	weatherApiMock := &WeatherApiMock{Results: weatherData}
	reader := &ReaderMock{DataSet: [][]byte{[]byte("some location"), []byte("1")}}
	stdoutMock := &StdoutMock{}

	cli := recommender.NewCli(geoApiMock, weatherApiMock, reader, stdoutMock.Println, stdoutMock.Printf)

	err := cli.Run()
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	recommendation := stdoutMock.Out[len(stdoutMock.Out)-1]

	if !strings.Contains(strings.ToLower(recommendation), "warm jacket") {
		t.Errorf("%s does not contain 'warm jacket'", recommendation)
	}
	if !strings.Contains(strings.ToLower(recommendation), "long pants") {
		t.Errorf("%s does not contain 'long pants'", recommendation)
	}
	if !strings.Contains(strings.ToLower(recommendation), "umbrella") {
		t.Errorf("%s does not contain 'umbrella'", recommendation)
	}
	if !strings.Contains(strings.ToLower(recommendation), "waterproof") {
		t.Errorf("%s does not contain 'waterproof'", recommendation)
	}
}

func TestCli(t *testing.T) {

	geoApiKey := os.Getenv("GEOAPIFY_API_KEY")
	if geoApiKey == "" {
		t.Skip("no GEOAPIFY_API_KEY set")
	}
	yrUserAgent := os.Getenv("YR_API_USER_AGENT")
	if yrUserAgent == "" {
		t.Skip("no YR_API_USER_AGENT set")
	}

	reader := &ReaderMock{DataSet: [][]byte{[]byte("some location"), []byte("1")}}
	stdoutMock := &StdoutMock{}

	cli := recommender.NewCli(
		api.NewGeoapifyClient(geoApiKey),
		api.NewWeatherApi(yrUserAgent),
		reader,
		stdoutMock.Println,
		stdoutMock.Printf)

	err := cli.Run()
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	recommendation := stdoutMock.Out[len(stdoutMock.Out)-1]

	if !strings.Contains(strings.ToLower(recommendation), "recommended clothing") {
		t.Errorf("%s does not contain 'recommended clothing'", recommendation)
	}
}
