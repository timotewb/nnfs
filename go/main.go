package main

import (
	"fmt"
	a "nnfs/app"
)

func main() {
	fmt.Println("main()")
	a.ASingleNeuron()
	a.ASingleNeuronNumpy()
	a.ALayerOfNeurons()
	a.ALayerOfNeuronsNumpy()
	a.BatchInputs()
	a.HiddenLayers()
	a.HiddenLayers_v2()
}
