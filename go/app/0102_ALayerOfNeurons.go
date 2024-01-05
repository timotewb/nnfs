package app

import "fmt"

func ALayerOfNeurons(){
	fmt.Println(" 01 ALayerOfNeurons()")
	var weights [3][4]float32
	var output float32
	var outputs [3]float32

	// 4 inputs
	inputs := [4]float32{1, 2, 3, 2.5}

	// 3 neurons
	weights[0] = [4]float32{0.2, 0.8, -0.5, 1}
	weights[1] = [4]float32{0.5, -0.91, 0.26, -0.5}
	weights[2] = [4]float32{-0.26, -0.27, 0.17, 0.87}
	bias := [3]float32{2, 3, .5}

	// compute output on 3 neurons
	// for a single layer
	// for each neuron
	for n:=0; n<len(weights); n++ {
		// for each input
		output = 0
		for i:=0; i<len(inputs); i++{
			output += inputs[i] * weights[n][i]
		}
		output += bias[n]
		outputs[n] = output
	}
	fmt.Println(" ",outputs)
}