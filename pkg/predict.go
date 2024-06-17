package pkg

import (
	"fmt"

	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
)

func PredictWithModel(stationMapTf map[string]*tf.Tensor) {
	model := tg.LoadModel("model/halim", []string{"serve"}, nil)

	fmt.Println(*model)

	result := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
		model.Op("StatefulPartitionedCall", 1),
	}, map[tf.Output]*tf.Tensor{
		model.Op("serving_default_input_2", 0): stationMapTf["Halim"],
	})

	floodResult := result[0].Value().([][]float32)
	climateResult := result[1].Value().([][]float32)

	fmt.Println("Flood_Branch result:", floodResult)
	fmt.Println("Climate_Branch result:", climateResult)
}
