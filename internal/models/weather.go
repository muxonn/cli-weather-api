package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Weather struct {
	Location struct {
		Name           string
		Country        string
		LocalTimeEpoch int64  `json:"localtime_epoch"`
		Timezone       string `json:"tz_id"`
	} `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

func (w *Weather) showBasicInformation() {

	loc, err := time.LoadLocation(w.Location.Timezone)

	if err != nil {
		fmt.Println("Error loading the timezone")
		return
	}

	currentTime := fmt.Sprintf("Current time: %s\n", time.Unix(w.Location.LocalTimeEpoch, 0).In(loc).Format("02-01-2006 15:04"))
	divider := ""

	for len(divider) < len(currentTime) {
		divider += "-"
	}
	fmt.Printf("%s\n", divider)
	fmt.Printf("City: %s\n", w.Location.Name)
	fmt.Printf("Country: %s\n", w.Location.Country)
	fmt.Printf(currentTime)
	fmt.Printf("%s\n", divider)
}

func (w *Weather) ShowCurrent() {
	w.showBasicInformation()
	fmt.Printf("Temperature: %.0f ℃\n", w.Current.TempC)
	fmt.Printf("Condition: %s\n", w.Current.Condition.Text)
}

func (w *Weather) ShowForecast() {
	todayHours := w.Forecast.ForecastDay[0].Hour
	tomorrowHours := w.Forecast.ForecastDay[1].Hour
	loc, err := time.LoadLocation(w.Location.Timezone)

	if err != nil {
		fmt.Println("Error loading the timezone")
		return
	}

	w.showBasicInformation()
	for _, hour := range todayHours {
		date := time.Unix(hour.TimeEpoch, 0).In(loc)
		if date.Before(time.Now().In(loc)) {
			continue
		}
		fmt.Printf("%s - %.0f ℃ | %s\n", date.Format("15:04"), hour.TempC, hour.Condition.Text)
	}

	fmt.Println("-----TOMORROW-----")

	for _, hour := range tomorrowHours {
		date := time.Unix(hour.TimeEpoch, 0).In(loc)
		if date.Hour() > time.Now().In(loc).Hour() {
			break
		}

		fmt.Printf("%s - %.0f ℃ | %s\n", date.Format("15:04"), hour.TempC, hour.Condition.Text)
	}

}

func (w *Weather) Decode(body []byte) error {
	err := json.Unmarshal(body, &w)

	if err != nil {
		return err
	}
	return nil
}
