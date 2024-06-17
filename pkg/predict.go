package pkg

import (
	"fmt"

	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
)

func PredictWithModel(stationMapTf map[string]*tf.Tensor) {
	model := tg.LoadModel("model/halim", []string{"serve"}, nil)

	fmt.Println(model)

	result := model.Exec([]tf.Output{
		model.Op("serving_default_input_1", 0),
	}, map[tf.Output]*tf.Tensor{
		model.Op("serving_default_input_1", 0): stationMapTf["Halim"],
	},
	)

	fmt.Println(result)
}
