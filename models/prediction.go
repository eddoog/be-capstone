package models

type LocationFloodPrediction struct {
	LocationName string  `json:"location_name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Value        float64 `json:"val"`
}
