package app

import "fmt"

// 3 neurons
type Layer1 struct {
	Weights [3][4]float32 `json:"weights"`
	Bias [3]float32 `json:"bias"`
}
// 3 neurons
type Layer2 struct {
	Weights [3][3]float32 `json:"weights"`
	Bias [3]float32 `json:"bias"`
}

func HiddenLayers(){
	fmt.Println(" 02 HiddenLayers()")

	inputs := [3][4]float32{{1,2,3,2.5},
	{2,5,-1,2},
	{-1.5,2.7,3.3,-.8}}

	var layer1 Layer1
	layer1.Weights[0] = [4]float32{0.2, 0.8, -0.5, 1}
	layer1.Weights[1] = [4]float32{0.5, -0.91, 0.26, -0.5}
	layer1.Weights[2] = [4]float32{-0.26, -0.27, 0.17, 0.87}

	layer1.Bias = [3]float32{2, 3, 0.5}

	var layer2 Layer2
	layer2.Weights[0] = [3]float32{0.1, -0.14, 0.5}
	layer2.Weights[1] = [3]float32{-0.5, 0.12, -0.33}
	layer2.Weights[2] = [3]float32{-0.44, 0.73, -0.13}

	layer2.Bias = [3]float32{-1, 2, -0.5}

}