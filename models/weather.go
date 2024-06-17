package models

import "encoding/json"

type Weather struct {
	Date string  `json:"date"` // Date in YYYY-MM-DD format
	Tavg float32 `json:"tavg"` // Average temperature in °C
	TMin float32 `json:"tmin"` // Minimum temperature in °C
	TMax float32 `json:"tmax"` // Maximum temperature in °C
	Prcp float32 `json:"prcp"` // Precipitation in mm
	Snow float32 `json:"snow"` // Snowfall in mm
	Wdir float32 `json:"wdir"` // Wind direction in degrees
	Wspd float32 `json:"wspd"` // Wind speed in km/h
	Wpgt float32 `json:"wpgt"` // Peak wind gust in km/h
	Pres float32 `json:"pres"` // Air pressure in hPa
	Tsun float32 `json:"tsun"` // Sunshine duration in minutes
}

type WeatherResponse struct {
	Meta struct {
		Generated string `json:"generated"`
	} `json:"meta"`
	Data []Weather `json:"data"`
}

func MarshalWeather(data []byte) ([]Weather, error) {
	var weatherResponse WeatherResponse
	err := json.Unmarshal(data, &weatherResponse)
	if err != nil {
		return nil, err
	}
	return weatherResponse.Data, nil
}
