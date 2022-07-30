package wheater

import "time"

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name      string    `json:"name"`
	Region    string    `json:"region"`
	Country   string    `json:"country"`
	Lat       float64   `json:"lat"`
	Lon       float64   `json:"lon"`
	TzId      string    `json:"tz_id"`
	LocalTime time.Time `json:"local_time"`
}

type Current struct {
	LastUpdateEpoch int64     `json:"last_update_epoch"`
	LastUpdated     string    `json:"last_updated"`
	TempC           float64   `json:"temp_c"`
	TempF           float64   `json:"temp_f"`
	IsDay           int8      `json:"is_day"`
	Condition       Condition `json:"condition"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int16  `json:"code"`
}
