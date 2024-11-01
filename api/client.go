package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"weather-api/config"
	"weather-api/internal/models"
)

func getWeatherApiUrl(location string, apiKey string) string {

	return fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=2&aqi=no&alerts=no", apiKey, location)
}

func GetWeather(location string, config config.Config) (*models.Weather, error) {
	resp, err := http.Get(getWeatherApiUrl(location, config.ApiKey))

	if err != nil {
		return nil, err
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
		return nil, err
	}

	var weather models.Weather

	err = weather.Decode(body)

	if err != nil {
		return nil, err
	}

	return &weather, nil
}
