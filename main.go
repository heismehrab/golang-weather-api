package main

import (
	"./utils/weather"
	"fmt"
)

func main() {
	data, err := weatherApi.GetCityWeather("tehran")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}