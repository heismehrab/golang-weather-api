package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/heismehrab/golang-weather-api/utils/weather"
	"os"
	"strings"
)

// Simply get a city from console input
// and get weather data from source.
func main() {
	fmt.Print("city:")
	reader := bufio.NewReader(os.Stdin)

	// Get city from console input.
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	if city == "" {
		panic("city is empty!")
	}

	// Get weather data.
	data := weatherApi.GetCityWeather(city)

	// Parsing the response.
	result := new(weatherApi.MainResponse)
	err := json.Unmarshal(data, &result)

	if err != nil {
		panic(err.Error())
	}

	// TODO. Handle response data with a beautiful readable structure.
	// TODO. I have to do something witch i dont remember now, at 2021-23-1 00:03.
	// TODO. Write a good document and explain the purpose of developing these codes.
}