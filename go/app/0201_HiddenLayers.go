package app

import "fmt"

// 3 neurons
type Layers struct {
	InputLayer   InputLayer
	HiddenLayers []HiddenLayer
	OutputLayer  OutputLayer
	layerOutput  []float32
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

	var layer2 HiddenLayer
	layer2.Weights[0] = [3]float32{0.1, -0.14, 0.5}
	layer2.Weights[1] = [3]float32{-0.5, 0.12, -0.33}
	layer2.Weights[2] = [3]float32{-0.44, 0.73, -0.13}

	layer2.Bias = [3]float32{-1, 2, -0.5}

	var layer3 HiddenLayer
	layer3.Weights[0] = [3]float32{0.1, -0.14, 0.5}
	layer3.Weights[1] = [3]float32{-0.5, 0.12, -0.33}
	layer3.Weights[2] = [3]float32{-0.44, 0.73, -0.13}

	layer3.Bias = [3]float32{-1, 2, -0.5}

	var layer4 OutputLayer
	layer4.Weights[0] = [3]float32{0.1, -0.14, 0.5}
	layer4.Weights[1] = [3]float32{-0.5, 0.12, -0.33}
	layer4.Weights[2] = [3]float32{-0.44, 0.73, -0.13}

	layer4.Bias = [3]float32{-1, 2, -0.5}

	var layers Layers
	layers.InputLayer = layer1
	layers.HiddenLayers = []HiddenLayer{layer2, layer3}
	layers.OutputLayer = layer4

	var n int
	var b int
	var h int
	var i int
	var v float32
	// for each batch (array in inputs) run through model and save
	for b = 0; b < len(inputs); b++ {
		//----------------------------------------------------------------------------------------
		// input layer
		//----------------------------------------------------------------------------------------
		layers.layerOutput = make([]float32, len(layers.InputLayer.Weights))
		// for each neuron
		for n = 0; n < len(layers.InputLayer.Bias); n++ {
			// for each input multiply by weight
			v = 0
			for i = 0; i < len(inputs[b]); i++ {
				v = v + inputs[b][i]*layers.InputLayer.Weights[n][i]
			}
			// add bias and append to output layer
			layers.layerOutput[n] = v + layers.InputLayer.Bias[n]
		}

		//----------------------------------------------------------------------------------------
		// hidden layers
		//----------------------------------------------------------------------------------------
		for h = 0; h < len(layers.HiddenLayers); h++ {
			inputs := layers.layerOutput
			layers.layerOutput = make([]float32, len(layers.HiddenLayers[h].Bias))
			// for each neuron
			for n = 0; n < len(layers.HiddenLayers[h].Bias); n++ {
				v = 0
				for i = 0; i < len(inputs); i++ {
					v = v + (inputs[i] * layers.HiddenLayers[h].Weights[n][i])
				}
				layers.layerOutput[n] = v + layers.HiddenLayers[h].Bias[n]
			}
		}

		//----------------------------------------------------------------------------------------
		// output layer
		//----------------------------------------------------------------------------------------
		batchOutput := make([]float32, len(layers.OutputLayer.Bias))
		for n = 0; n < len(layers.OutputLayer.Weights); n++ {
			v = 0
			for i = 0; i < len(layers.layerOutput); i++ {
				v = v + (layers.layerOutput[i] * layers.OutputLayer.Weights[n][i])
			}
			batchOutput[n] = v + layers.OutputLayer.Bias[n]
		}
		fmt.Println(" ", batchOutput)
	}

	// we run one row of input through the entire model and save to array, then run next input

}
