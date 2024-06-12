package models

import "encoding/json"

type Weather struct {
	Date string  `json:"date"` // Date in YYYY-MM-DD format
	Tavg float64 `json:"tavg"` // Average temperature in °C
	TMin float64 `json:"tmin"` // Minimum temperature in °C
	TMax float64 `json:"tmax"` // Maximum temperature in °C
	Prcp float64 `json:"prcp"` // Precipitation in mm
	Snow float64 `json:"snow"` // Snowfall in mm
	Wdir float64 `json:"wdir"` // Wind direction in degrees
	Wspd float64 `json:"wspd"` // Wind speed in km/h
	Wpgt float64 `json:"wpgt"` // Peak wind gust in km/h
	Pres float64 `json:"pres"` // Air pressure in hPa
	Tsun float64 `json:"tsun"` // Sunshine duration in minutes
}

func MarshalWeather(data []byte) ([]Weather, error) {
	var weather []Weather
	err := json.Unmarshal(data, &weather)
	if err != nil {
		return nil, err
	}
	return weather, nil
}
