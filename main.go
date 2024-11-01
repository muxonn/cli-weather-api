package main

import (
	"weather-api/api"
	"weather-api/cli"
	"weather-api/config"
	"weather-api/internal/types"
)

func main() {

	args := cli.ParseCliArgs()

	cfg := config.LoadConfig()

	weather, err := api.GetWeather(args.Location, cfg)

	if err != nil {
		panic(err)
	}

	if args.WeatherType == types.Forecast {
		weather.ShowForecast()
		return
	}
	weather.ShowCurrent()

}
