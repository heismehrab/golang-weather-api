package weatherApi

import (
	"fmt"
	"github.com/heismehrab/golang-weather-api/config"
	"io/ioutil"
	"net/http"
)

// Get url data from config file.
func GetApiUrl() string {
	return config.CurrentWeatherApiUrl + config.ApiKey
}

// Send a request and get
// weather data from stations.
func GetCityWeather(city string) []byte {
	requestUrl :=  fmt.Sprintf(GetApiUrl(), city)
	res, err := http.Get(requestUrl)

	if err != nil {
		panic(err.Error())
	}

	result, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	return result
}
