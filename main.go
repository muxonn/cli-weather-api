package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getUrl(location string) string {
	apiKey := os.Getenv("CLI_WEATHER_API_KEY")
	return fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=2&aqi=no&alerts=no", apiKey, location)
}

type Forecast struct {
	ForecastDay []struct {
		Hour []struct {
			TimeEpoch int64   `json:"time_epoch"`
			TempC     float64 `json:"temp_c"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
		} `json:"hour"`
	} `json:"forecastday"`
}

type Current struct {
	TempC     float64 `json:"temp_c"`
	Condition struct {
		Text string `json:"text"`
	} `json:"condition"`
}

type Weather struct {
	Location struct {
		Name           string
		Country        string
		LocalTimeEpoch int64 `json:"localtime_epoch"`
	} `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

func (w *Weather) showCurrentWeather() {
	fmt.Printf("City: %s\n", w.Location.Name)
	fmt.Printf("Country: %s\n", w.Location.Country)
	fmt.Printf("Current time: %s\n", time.Unix(w.Location.LocalTimeEpoch, 0).Format("02-01-2006 15:04"))
	fmt.Printf("Temperature: %.0f ℃\n", w.Current.TempC)
	fmt.Printf("Condition: %s\n", w.Current.Condition.Text)
}

func (w *Weather) showForecast() {
	hours := w.Forecast.ForecastDay[0].Hour

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}
		fmt.Printf("%s - %.0f ℃ | %s\n", date.Format("15:04"), hour.TempC, hour.Condition.Text)
	}
}

func (w *Weather) decode(body []byte) error {
	err := json.Unmarshal(body, &w)

	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	resp, err := http.Get(getUrl("London"))

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather

	err = weather.decode(body)

	if err != nil {
		panic(err)
	}

	weather.showCurrentWeather()
	weather.showForecast()
}
