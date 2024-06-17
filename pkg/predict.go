package pkg

import (
	"fmt"
	"strings"
	"sync"

	"github.com/eddoog/be-capstone/models"
	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
)

func PredictWithModel(stationMapTf map[string]*tf.Tensor) {
	var wg sync.WaitGroup

	predictionsResult := make(chan []models.LocationFloodPrediction)
	errCh := make(chan error, len(stationMapTf))

	for stationName, stationTf := range stationMapTf {
		wg.Add(1)

		go callModel(stationName, stationTf, &wg, predictionsResult, errCh)
	}

	wg.Wait()

	close(predictionsResult)
	close(errCh)
}

func callModel(
	stationName string,
	stationTf *tf.Tensor,
	wg *sync.WaitGroup,
	predictionResult chan<- []models.LocationFloodPrediction,
	errCh chan<- error,
) {
	defer wg.Done()
	defer SendInfoLog("Finished Predicting data for station " + stationName)

	modelOutputName := GetOutputModelName()
	model := tg.LoadModel(fmt.Sprintf("model/%s", strings.ToLower(stationName)), []string{"serve"}, nil)

	result := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
		model.Op("StatefulPartitionedCall", 1),
	}, map[tf.Output]*tf.Tensor{
		model.Op(modelOutputName[stationName], 0): stationTf,
	})

	climateResult := result[0].Value().([][]float32)
	floodResult := result[1].Value().([][]float32)

	fmt.Println(stationName)
	fmt.Println(climateResult)
	fmt.Println(floodResult)
}

/* model := tg.LoadModel("model/halim", []string{"serve"}, nil)

result := model.Exec([]tf.Output{
	model.Op("StatefulPartitionedCall", 0),
	model.Op("StatefulPartitionedCall", 1),
}, map[tf.Output]*tf.Tensor{
	model.Op("serving_default_input_2", 0): stationMapTf["Halim"],
})

climateResult := result[0].Value().([][]float32)
floodResult := result[1].Value().([][]float32) */
