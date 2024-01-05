package app

import "fmt"

func BatchInputs(){
	fmt.Println(" 01 BatchInputs()")

	//declare vars
	var inputs [][4]float32
	var weights [][4]float32
	var outputs [][3]float32
	var output [3]float32
	var val float32

	//init vars
	inputs = append(inputs, [4]float32{1,2,3,2.5})
	inputs = append(inputs, [4]float32{2,5,-1,2})
	inputs = append(inputs, [4]float32{-1.5,2.7,3.3,-.8})

	weights = append(weights, [4]float32{0.2, 0.8, -0.5, 1})
	weights = append(weights, [4]float32{0.5, -0.91, 0.26, -0.5})
	weights = append(weights, [4]float32{-0.26, -0.27, 0.17, 0.87})

	bias := [3]float32{2,3,.5}

	// for each input batch
	for b:=0; b <len(inputs); b++{
		// for each neuron
		output = [3]float32{}
		for n:=0; n<len(weights) ; n++{
			// for each input
			val = 0
			for i:=0 ; i<len(inputs[b]) ; i++{
				val = val + weights[n][i] * inputs[b][i]
			}
			output[n] = val + bias[n]
		}
		outputs = append(outputs, output)
	}
	fmt.Println(" ", outputs)
}

// need to output a matrix of output batches