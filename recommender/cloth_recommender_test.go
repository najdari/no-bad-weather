package recommender_test

import (
	"testing"

	"nimaajdari.com/nobadweather/api"
	"nimaajdari.com/nobadweather/recommender"
)

func TestClothRecommender(t *testing.T) {
	cases := []struct {
		Name     string
		Expected string
		In       []api.WeatherData
	}{
		{
			Name:     "Very cold and dry",
			Expected: "Recommended clothing: Heavy winter coat, gloves, and warm hat. ",
			In: []api.WeatherData{
				{
					AirTemperature:      -10,
					WindSpeed:           4.5,
					PrecipitationAmount: 0,
				},
			},
		},
		{
			Name:     "Cold and dry",
			Expected: "Recommended clothing: Warm jacket, long pants. ",
			In: []api.WeatherData{
				{
					AirTemperature:      1,
					WindSpeed:           4.5,
					PrecipitationAmount: 0,
				},
			},
		},
		{
			Name:     "Mild and dry",
			Expected: "Recommended clothing: Light jacket or sweater. ",
			In: []api.WeatherData{
				{
					AirTemperature:      10,
					WindSpeed:           4.5,
					PrecipitationAmount: 0,
				},
			},
		},
		{
			Name:     "Warm and dry",
			Expected: "Recommended clothing: T-shirt and shorts. ",
			In: []api.WeatherData{
				{
					AirTemperature:      20,
					WindSpeed:           4.5,
					PrecipitationAmount: 0,
				},
			},
		},
		{
			Name:     "Very cold with precipitation",
			Expected: "Recommended clothing: Heavy winter coat, gloves, and warm hat. Bring an umbrella or wear waterproof clothing. ",
			In: []api.WeatherData{
				{
					AirTemperature:      -10,
					WindSpeed:           4.5,
					PrecipitationAmount: 5,
				},
			},
		},
		{
			Name:     "Cold with precipitation",
			Expected: "Recommended clothing: Warm jacket, long pants. Bring an umbrella or wear waterproof clothing. ",
			In: []api.WeatherData{
				{
					AirTemperature:      1,
					WindSpeed:           4.5,
					PrecipitationAmount: 5,
				},
			},
		},
		{
			Name:     "Mild with precipitation",
			Expected: "Recommended clothing: Light jacket or sweater. Bring an umbrella or wear waterproof clothing. ",
			In: []api.WeatherData{
				{
					AirTemperature:      10,
					WindSpeed:           4.5,
					PrecipitationAmount: 5,
				},
			},
		},
		{
			Name:     "Warm with precipitation",
			Expected: "Recommended clothing: T-shirt and shorts. Bring an umbrella or wear waterproof clothing. ",
			In: []api.WeatherData{
				{
					AirTemperature:      20,
					WindSpeed:           4.5,
					PrecipitationAmount: 5,
				},
			},
		},
		{
			Name:     "Very cold with precipitation and wind",
			Expected: "Recommended clothing: Heavy winter coat, gloves, and warm hat. Bring an umbrella or wear waterproof clothing. Consider wearing a windbreaker. ",
			In: []api.WeatherData{
				{
					AirTemperature:      -10,
					WindSpeed:           8,
					PrecipitationAmount: 5,
				},
			},
		},
		{
			Name:     "Cold with precipitation and wind",
			Expected: "Recommended clothing: Warm jacket, long pants. Bring an umbrella or wear waterproof clothing. Consider wearing a windbreaker. ",
			In: []api.WeatherData{
				{
					AirTemperature:      1,
					WindSpeed:           8,
					PrecipitationAmount: 5,
				},
			},
		},
		{
			Name:     "Mild with precipitation and wind",
			Expected: "Recommended clothing: Light jacket or sweater. Bring an umbrella or wear waterproof clothing. Consider wearing a windbreaker. ",
			In: []api.WeatherData{
				{
					AirTemperature:      10,
					WindSpeed:           8,
					PrecipitationAmount: 5,
				},
			},
		},
		{
			Name:     "Warm with precipitation and wind",
			Expected: "Recommended clothing: T-shirt and shorts. Bring an umbrella or wear waterproof clothing. Consider wearing a windbreaker. ",
			In: []api.WeatherData{
				{
					AirTemperature:      20,
					WindSpeed:           8,
					PrecipitationAmount: 5,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			out, err := recommender.RecommendClothing(c.In)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if out != c.Expected {
				t.Errorf("Expected %s, got %s", c.Expected, out)
			}
		})
	}

}
