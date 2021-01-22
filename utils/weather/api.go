package weatherApi

import (
	"../../config"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetApiUrl() string {
	return config.CurrentWeatherApiUrl + config.ApiKey
}

func GetCityWeather(city string) ([]byte, error) {

	if city == "" {
		return make([]byte, 1), errors.New("argument city is empty")
	}

	requestUrl :=  strings.Replace(GetApiUrl(), "{city}", city, 1)
	res, err := http.Get(requestUrl)

	if err != nil {
		return make([]byte, 1), err
	}

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return make([]byte, 1), err
	}

	return resBody, err
}
