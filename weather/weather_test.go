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

func TestGetWeatherIncorrectFormat(t *testing.T) {
	geo := geo.GeoData{
		City: "London",
	}
	format := 125
	expected := weather.ErrorIncorrectFormat

	_, err := weather.GetWeather(geo, format)
	if err != expected {
		t.Errorf("Ожидалось %v, получили %v", expected, err)
	}
}
