package pkg

import (
	"github.com/eddoog/be-capstone/models"
)

var (
	stasionsInstance []*models.Station
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
	once.Do(InitializeStation)
	return stasionsInstance
}
