package app

import "fmt"

// 3 neurons
type Layers struct {
	InputLayer   InputLayer
	HiddenLayers []HiddenLayer
	OutputLayer  OutputLayer
}

// 3 neurons
type InputLayer struct {
	Weights [3][4]float32 `json:"weights"`
	Bias    [3]float32    `json:"bias"`
}

// 3 neurons
type HiddenLayer struct {
	Weights [3][3]float32 `json:"weights"`
	Bias    [3]float32    `json:"bias"`
}

// 3 neurons
type OutputLayer struct {
	Weights [3][3]float32 `json:"weights"`
	Bias    [3]float32    `json:"bias"`
}

func HiddenLayers() {
	fmt.Println(" 02 HiddenLayers()")

	inputs := [3][4]float32{{1, 2, 3, 2.5},
		{2, 5, -1, 2},
		{-1.5, 2.7, 3.3, -.8}}

	var layer1 InputLayer
	layer1.Weights[0] = [4]float32{0.2, 0.8, -0.5, 1}
	layer1.Weights[1] = [4]float32{0.5, -0.91, 0.26, -0.5}
	layer1.Weights[2] = [4]float32{-0.26, -0.27, 0.17, 0.87}

	layer1.Bias = [3]float32{2, 3, 0.5}

	// no hidden layer

	var layer2 OutputLayer
	layer2.Weights[0] = [3]float32{0.1, -0.14, 0.5}
	layer2.Weights[1] = [3]float32{-0.5, 0.12, -0.33}
	layer2.Weights[2] = [3]float32{-0.44, 0.73, -0.13}

	layer2.Bias = [3]float32{-1, 2, -0.5}

	var layers Layers
	layers.InputLayer = layer1
	layers.OutputLayer = layer2

	var n int
	var w int
	var b int
	var h int
	var v float32
	// for each batch (array in inputs) run through model and save
	for b = 0; b < len(inputs); b++ {
		fmt.Println("\nBatch:", inputs[b])
		//----------------------------------------------------------------------------------------
		// input layer
		//----------------------------------------------------------------------------------------
		layerOutput := make([]float32, len(layers.InputLayer.Weights))
		// for each neuron
		for n = 0; n < len(layers.InputLayer.Weights); n++ {
			// for each weight/input pair
			v = 0
			for w = 0; w < len(layers.InputLayer.Weights[n]); w++ {
				v = v + (inputs[b][n] * layers.InputLayer.Weights[n][w])
			}
			layerOutput[n] = v
		}

		//----------------------------------------------------------------------------------------
		// hidden layers
		//----------------------------------------------------------------------------------------
		for h = 0; h < len(layers.HiddenLayers); h++ {
			inputs := layerOutput
			layerOutput := make([]float32, len(layers.HiddenLayers))
			// for each neuron
			for n = 0; n < len(layers.HiddenLayers[h].Weights); n++ {
				v = 0
				for w = 0; w < len(layers.HiddenLayers[h].Weights[n]); w++ {
					v = v + (inputs[w] * layers.HiddenLayers[h].Weights[n][w])
				}
				layerOutput[n] = v
			}
		}

		//----------------------------------------------------------------------------------------
		// output layer
		//----------------------------------------------------------------------------------------
		batchOutput := make([]float32, len(layers.OutputLayer.Weights))
		for n = 0; n < len(layers.OutputLayer.Weights); n++ {
			v = 0
			for w = 0; w < len(layers.OutputLayer.Weights[n]); w++ {
				v = v + (layerOutput[w] * layers.OutputLayer.Weights[n][w])
			}
			batchOutput[n] = v
		}
		fmt.Println(layerOutput)
	}

	// we run one row of input through the entire model and save to array, then run next input

}
