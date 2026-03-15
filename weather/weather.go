package weather

import (
	"demo/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.GeoData, format int) (string, error) {
	baseUrl, err := url.Parse("https://www.wttr.in/" + geo.City)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("city", geo.City)
	params.Set("format", fmt.Sprint(format))

	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("NOT_200")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
