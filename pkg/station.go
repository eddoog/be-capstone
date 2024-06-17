package pkg

import (
	"sync"

	"github.com/eddoog/be-capstone/models"
)

var (
	stasionsInstance []*models.Station
	stationOnce      sync.Once
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

func GetStations() []*models.Station {
	stationOnce.Do(InitializeStation)
	return stasionsInstance
}
