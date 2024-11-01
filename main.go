package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"weather-api/internal/models"
)

func getUrl(location string) string {
	apiKey := os.Getenv("CLI_WEATHER_API_KEY")
	return fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=2&aqi=no&alerts=no", apiKey, location)
}

func main() {
	err := godotenv.Load()
	args := os.Args

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(args) < 2 {
		log.Fatal("Not enough arguments, you have to provide the location")
	}

	arg := args[1]

	resp, err := http.Get(getUrl(arg))

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("Something went wrong...")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var weather models.Weather

	err = weather.Decode(body)

	if err != nil {
		panic(err)
	}

	if len(args) > 2 && args[2] == "-f" {
		weather.ShowForecast()
		return
	}
	weather.ShowCurrent()

}
