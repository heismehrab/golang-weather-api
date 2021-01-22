package main

import (
	"./utils/weather"
	"bufio"
	"encoding/json"
	"fmt"
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
}