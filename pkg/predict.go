package pkg

import (
	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
)

func PredictWithModel(stationMapTf map[string]*tf.Tensor) {
	model := tg.LoadModel("model/halim", []string{"serve"}, nil)

	result := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
		model.Op("StatefulPartitionedCall", 1),
	}, map[tf.Output]*tf.Tensor{
		model.Op("serving_default_input_2", 0): stationMapTf["Halim"],
	})

	climateResult := result[0].Value().([][]float32)
	floodResult := result[1].Value().([][]float32)
}
