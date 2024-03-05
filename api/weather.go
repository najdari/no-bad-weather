package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"nimaajdari.com/nobadweather/errs"
)

type WeatherResponse struct {
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Properties struct {
	Meta       Meta        `json:"meta"`
	Timeseries []Timeserie `json:"timeseries"`
}

type Meta struct {
	UpdatedAt string            `json:"updated_at"`
	Units     map[string]string `json:"units"`
}

type Timeserie struct {
	Time string        `json:"time"`
	Data TimeserieData `json:"data"`
}

type TimeserieData struct {
	Instant     Instant       `json:"instant"`
	Next12Hours SummaryDetail `json:"next_12_hours"`
	Next1Hour   SummaryDetail `json:"next_1_hours"`
	Next6Hours  SummaryDetail `json:"next_6_hours"`
}

type Instant struct {
	Details InstantDetails `json:"details"`
}

type InstantDetails struct {
	AirPressureAtSeaLevel float64 `json:"air_pressure_at_sea_level"`
	AirTemperature        float64 `json:"air_temperature"`
	CloudAreaFraction     float64 `json:"cloud_area_fraction"`
	RelativeHumidity      float64 `json:"relative_humidity"`
	WindFromDirection     float64 `json:"wind_from_direction"`
	WindSpeed             float64 `json:"wind_speed"`
}

type SummaryDetail struct {
	Details SummaryDetails `json:"details,omitempty"`
	Summary struct {
		SymbolCode string `json:"symbol_code"`
	} `json:"summary"`
}

type SummaryDetails struct {
	PrecipitationAmount float64 `json:"precipitation_amount"`
}

type WeatherData struct {
	AirTemperature      float64
	CloudAreaFraction   float64
	RelativeHumidity    float64
	WindSpeed           float64
	PrecipitationAmount float64
}

type WeatherApi struct {
	userAgent string
}

func NewWeatherApi(userAgnet string) WeatherApi {
	return WeatherApi{userAgent: userAgnet}
}

func (w WeatherApi) GetWeatherInfo(lat, lon float64) ([]WeatherData, error) {

	url := fmt.Sprintf("https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=%f&lon=%f", lat, lon)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errs.StackInfo(err)
	}

	req.Header.Set("User-Agent", w.userAgent)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errs.StackInfo(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errs.Wrap("error reading response body", err)
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return nil, errs.Wrap("error unmarshing weather api response", err)
	}

	weatherDataList := make([]WeatherData, 0, len(weatherResponse.Properties.Timeseries))
	for _, v := range weatherResponse.Properties.Timeseries {
		weatherData := WeatherData{
			AirTemperature:      v.Data.Instant.Details.AirTemperature,
			CloudAreaFraction:   v.Data.Instant.Details.CloudAreaFraction,
			RelativeHumidity:    v.Data.Instant.Details.RelativeHumidity,
			WindSpeed:           v.Data.Instant.Details.WindSpeed,
			PrecipitationAmount: v.Data.Next1Hour.Details.PrecipitationAmount,
		}
		weatherDataList = append(weatherDataList, weatherData)
	}

	return weatherDataList, nil
}
