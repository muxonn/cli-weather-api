package types

import "errors"

type WeatherType int

const (
	Current WeatherType = iota
	Forecast
)

var weatherTypeName = map[WeatherType]string{
	Current:  "-c",
	Forecast: "-f",
}

func (wt WeatherType) String() string {
	return weatherTypeName[wt]
}

func ParseWeatherType(flag string) (WeatherType, error) {
	switch flag {
	case "-c":
		return Current, nil
	case "-f":
		return Forecast, nil
	default:
		return 0, errors.New("invalid weather type: use -c for current or -f for forecast")
	}
}
