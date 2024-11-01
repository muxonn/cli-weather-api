package models

type Current struct {
	TempC     float64 `json:"temp_c"`
	Condition struct {
		Text string `json:"text"`
	} `json:"condition"`
}
