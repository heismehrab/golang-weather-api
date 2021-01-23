package weatherApi

import "fmt"

type MainResponse struct {
	Coord coord `json:"coord"`
	Weather []weather `json:"weather"`
	Base string `json:"base"`
	Main main `json:"main"`
	Visibility int `json:"visibility"`
	Wind wind `json:"wind"`
	Timezone int `json:"timezone"`
	Id int `json:"id"`
	Name string `json:"name"`
	Cod int `json:"cod"`
}

type coord struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type weather struct {
	Id int `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type main struct {
	Temp float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	FempMin float32 `json:"femp_min"`
	TtempMax float32 `json:"temp_max"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"`
	Clouds clouds `json:"clouds"`
	Dt int `json:"dt"`
	Sys sys `json:"sys"`
}

type wind struct {
	Speed float32 `json:"speed"`
	Deg int `json:"deg"`
}

type clouds struct {
	All int `json:"all"`
}

type sys struct {
	Type int `json:"type"`
	Id int `json:"id"`
	Country string `json:"country"`
	Sunrise int `json:"sunrise"`
	Sunset int `json:"sunset"`
}

func CliOutput(response *MainResponse) string {
	output := " city : %s (%s,%s) \n"
	output = output + ""


	return fmt.Sprintf(
		output,
		response.Name,
		response.Coord.Lat,
		response.Coord.Lon,
		)
}
