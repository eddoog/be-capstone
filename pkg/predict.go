package pkg

import (
	"fmt"
	"strings"
	"sync"

	"github.com/eddoog/be-capstone/models"
	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
)

func PredictWithModel(stationMapTf map[string]*tf.Tensor) ([]models.LocationFloodPrediction, error) {
	var wg sync.WaitGroup

	predictionsResult := make(chan []models.LocationFloodPrediction, len(stationMapTf))
	errCh := make(chan error, len(stationMapTf))

	for stationName, stationTf := range stationMapTf {
		wg.Add(1)

		go callModel(stationName, stationTf, &wg, predictionsResult, errCh)
	}

	wg.Wait()

	close(predictionsResult)
	close(errCh)

	result := []models.LocationFloodPrediction{}

	for predictions := range predictionsResult {
		result = append(result, predictions...)
	}

	if len(errCh) > 0 {
		return nil, <-errCh
	}

	return result, nil
}

func callModel(
	stationName string,
	stationTf *tf.Tensor,
	wg *sync.WaitGroup,
	ch chan<- []models.LocationFloodPrediction,
	errCh chan<- error,
) {
	defer func() {
		if r := recover(); r != nil {
			SendWarnLog("Error predicting data for station " + stationName + ": " + r.(error).Error())
			errCh <- r.(error)
		} else {
			SendInfoLog("Finished Predicting data for station " + stationName)
		}
		wg.Done()
	}()

	modelOutputName := GetOutputModelName()
	model := tg.LoadModel(fmt.Sprintf("model/%s", strings.ToLower(stationName)), []string{"serve"}, nil)

	result := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
		model.Op("StatefulPartitionedCall", 1),
	}, map[tf.Output]*tf.Tensor{
		model.Op(modelOutputName[stationName], 0): stationTf,
	})

	floodResult := result[1].Value().([][]float32)

	locations := GetPlaces()[stationName]

	var predictions []models.LocationFloodPrediction

	for i, location := range locations {
		predictions = append(predictions, models.LocationFloodPrediction{
			LocationName: location.LocationName,
			Latitude:     location.Latitude,
			Longitude:    location.Longitude,
			Value:        float64(floodResult[0][i]),
			Color:        DecideHexColor(float64(floodResult[0][i])),
		})
	}

	SendInfoLog("Predictions for station " + stationName + " : " + fmt.Sprint(predictions))

	ch <- predictions
}
