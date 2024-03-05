package recommender

import (
	"errors"
	"fmt"

	"nimaajdari.com/nobadweather/api"
)

func RecommendClothing(weatherData []api.WeatherData) (string, error) {

	if len(weatherData) == 0 {
		return "", errors.New("no weather data available for cloth recommendation")
	}

	currentWeather := weatherData[0]

	fmt.Printf("Recommending clothes for weather data: %+v\n", currentWeather)

	temperature := currentWeather.AirTemperature
	precipitation := currentWeather.PrecipitationAmount
	wind := currentWeather.WindSpeed

	recommendation := "Recommended clothing: "

	if temperature < -5 {
		recommendation += "Heavy winter coat, gloves, and warm hat. "
	} else if temperature < 5 {
		recommendation += "Warm jacket, long pants. "
	} else if temperature < 15 {
		recommendation += "Light jacket or sweater. "
	} else {
		recommendation += "T-shirt and shorts. "
	}

	if precipitation > 0 {
		recommendation += "Bring an umbrella or wear waterproof clothing. "
	}

	if wind > 5 {
		recommendation += "Consider wearing a windbreaker. "
	}

	return recommendation, nil
}
