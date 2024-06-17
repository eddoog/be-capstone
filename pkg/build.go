package pkg

import (
	"github.com/eddoog/be-capstone/models"
	tf "github.com/galeone/tensorflow/tensorflow/go"
)

func BuildStationMapTf(weatherMap map[string][]models.Weather) (map[string]*tf.Tensor, error) {
	// BuildStationMapTf is a function to build a map of stations
	// with the key is the station name and the value is the tensorflow array

	tfMapArray := make(map[string]*tf.Tensor)

	for stationName, weathers := range weatherMap {
		tfArray := make([][]float64, len(weathers))

		for idx, weather := range weathers {
			convertedWeather, err := buildTfArray(weather)

			if err != nil {
				SendWarnLog(err.Error())
				return nil, err
			}

			tfArray[idx] = convertedWeather.Value().([]float64)
		}

		threeDimArray := [][][]float64{tfArray}
		tensor, err := tf.NewTensor(threeDimArray)
		if err != nil {
			SendWarnLog(err.Error())
			return nil, err
		}

		tfMapArray[stationName] = tensor
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
