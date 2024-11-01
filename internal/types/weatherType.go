package types

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

func ParseWeatherType(flag string) WeatherType {
	switch flag {
	case "-c":
		return Current
	case "-f":
		return Forecast
	default:
		return Current
	}
}
