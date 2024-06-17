package pkg

import (
	"github.com/eddoog/be-capstone/models"
	tf "github.com/galeone/tensorflow/tensorflow/go"
)

func BuildStationMapTf(weatherMap map[string][]models.Weather) (map[string][][]float64, error) {
	// BuildStationMapTf is a function to build a map of stations
	// with the key is the station name and the value is the tensorflow array

	tfMapArray := make(map[string][][]float64)

	for stationName, weathers := range weatherMap {
		tfArray := make([][]float64, len(weathers))

		for _, weather := range weathers {
			convertedWeather, err := buildTfArray(weather)

			if err != nil {
				return nil, err
			}

			tfArray = append(tfArray, convertedWeather.Value().([]float64))
		}

		tfMapArray[stationName] = tfArray
	}

	return tfMapArray, nil
}

func buildTfArray(weather models.Weather) (*tf.Tensor, error) {
	return tf.NewTensor([]float64{
		weather.Prcp,
		weather.Tavg,
		weather.Wspd,
	})
}
