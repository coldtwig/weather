package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected result, данные для функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error("Ошибка получения города:" + err.Error())
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось: %v, получено: %v", expected, got)
	}
	// Assert - проверка результата с expected
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londonandsd"

	_, err := geo.GetMyLocation(city)
	if err != geo.ErrorIncorrectCity {
		t.Errorf("Ожидалась ошибка: %v, получено: %v", geo.ErrorIncorrectCity, err)
	}
}
