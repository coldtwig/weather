package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type cityPopulationResp struct {
	Error bool `json:"error"`
}

var ErrorIncorrectCity = errors.New("INCORRECT_CITY")
var ErrorNot200 = errors.New("NOT_200")

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity, err := checkLocation(city)
		if err != nil {
			return nil, err
		}
		if !isCity {
			return nil, ErrorIncorrectCity
		}

		return &GeoData{
			City: city,
		}, nil
	}

	resp, err := http.Get("https://ipapi.co/json")
	if err != nil {
		return nil, err
	}
	if (*resp).StatusCode != 200 {
		return nil, ErrorNot200
	}
	defer resp.Body.Close()

	var geo GeoData
	err = json.NewDecoder(resp.Body).Decode(&geo)
	if err != nil {
		return nil, err
	}

	return &geo, nil
}

func checkLocation(city string) (bool, error) {
	if city == "" {
		return false, nil
	}

	postBody, err := json.Marshal(GeoData{
		City: city,
	})
	if err != nil {
		return false, err
	}

	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false, err
	}
	if resp.StatusCode != 200 {
		return false, errors.New("STATUS_NOT_201")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var populationResp cityPopulationResp
	err = json.Unmarshal(body, &populationResp)
	if err != nil {
		return false, err
	}

	return !populationResp.Error, nil
}
