package models

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
