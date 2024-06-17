package pkg

import (
	"sync"

	"github.com/eddoog/be-capstone/models"
)

var (
	stasionsInstance []*models.Station
	outputModelName  map[string]string
	stationOnce      sync.Once
	outputNameOnce   sync.Once
)

func InitializeStation() {
	stasionsInstance = []*models.Station{
		{
			StationName: "Kemayoran",
			StationId:   96749,
		},
		{
			StationName: "TanjungPriok",
			StationId:   96741,
		},
		{
			StationName: "Halim",
			StationId:   96747,
		},
	}
}

func InitializeModelOutputName() {
	outputModelName = map[string]string{
		"Kemayoran":    "serving_default_input_6",
		"TanjungPriok": "serving_default_input_1",
		"Halim":        "serving_default_input_2",
	}
}

func GetStations() []*models.Station {
	stationOnce.Do(InitializeStation)
	return stasionsInstance
}

func GetOutputModelName() map[string]string {
	outputNameOnce.Do(InitializeModelOutputName)
	return outputModelName
}
