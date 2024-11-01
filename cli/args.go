package cli

import (
	"log"
	"os"
	"weather-api/internal/types"
)

type WeatherArgs struct {
	Location    string
	WeatherType types.WeatherType
}

func ParseCliArgs() *WeatherArgs {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Not enough arguments, you have to provide the location")
	}

	location := args[1]

	weatherType := types.Current
	if len(args) > 2 {
		weatherType = types.ParseWeatherType(args[2])
	}

	return &WeatherArgs{Location: location, WeatherType: weatherType}
}
