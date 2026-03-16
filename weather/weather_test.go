package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "London"
	geo := geo.GeoData{
		City: expected,
	}
	format := 125

	currentResult, err := weather.GetWeather(geo, format)
	if err != nil {
		t.Errorf("Ожидалось %v, получили %v", expected, err)
	}
	if !strings.Contains(currentResult, expected) {
		t.Errorf("Ожидалось %v, получили %v", expected, currentResult)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{
		name:   "Big format",
		format: 147,
	},
	{
		name:   "0 format",
		format: 0,
	},
	{
		name:   "Minus format",
		format: -1,
	},
}

func TestGetWeatherIncorrectFormat(t *testing.T) {
	geo := geo.GeoData{
		City: "London",
	}
	expected := weather.ErrorIncorrectFormat

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := weather.GetWeather(geo, testCase.format)
			if err != expected {
				t.Errorf("Ожидалось %v, получили %v", expected, err)
			}
		})
	}
}
