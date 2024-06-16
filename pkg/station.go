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
			StationName: "Jakarta / Soekarno-Hatta",
			StationId:   96749,
		},
		{
			StationName: "Jakarta / Tanjung Priok",
			StationId:   96741,
		},
		{
			StationName: "Jakarta / Halim Perdana Kusuma",
			StationId:   96747,
		},
	}
}

func GetStations() []*models.Station {
	stationOnce.Do(InitializeStation)
	return stasionsInstance
}
