package main

import (
	"fmt"
	a "nnfs/app"
)

func main(){
	fmt.Println("main()")
	a.ASingleNeuron()
	a.ALayerOfNeurons()
	a.BatchInputs()
	a.HiddenLayers()
}