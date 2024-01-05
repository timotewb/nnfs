package app

import "fmt"

func ASingleNeuron(){
	fmt.Println(" 01 ASingleNeuron()")
	var inputs []float32
	var weights []float32
	var bias float32
	var outputs float32

	inputs = append(inputs, 1, 2, 3)
	weights = append(weights, .2, .8, -.5)
	bias = 2

	for i:=0 ; i<len(inputs); i++ {
		outputs += inputs[i]*weights[i]
	}
	outputs += bias

	fmt.Println(" ",outputs)
}