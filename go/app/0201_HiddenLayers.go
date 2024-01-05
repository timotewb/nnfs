package app

import "fmt"

func HiddenLayers(){
	fmt.Println(" 02 HiddenLayers()")
	inputs := [3][4]float32{{1,2,3,2.5},
	{2,5,-1,2},
	{-1.5,2.7,3.3,-.8}}
	weights := [][]float32{{0.2, 0.8, -0.5, 1},
	{0.5, -0.91, 0.26, -0.5},
	{-0.26, -0.27, 0.17, 0.87}}

	fmt.Println(inputs)
	fmt.Println(weights)
}