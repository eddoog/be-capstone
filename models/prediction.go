package models

type LocationFloodPrediction struct {
	Location
	Value float64 `json:"val"`
}
