package app

import (
	"fmt"
	"log"

	"github.com/timotewb/gonn/numpy"
)

func ALayerOfNeuronsNumpy() {
	fmt.Println(" 01 ALayerOfNeuronsNumpy()")

	inputs := []float64{1, 2, 3, 2.5}
	weights := [][]float64{{0.2, 0.8, -0.5, 1}, {0.5, -0.91, 0.26, -0.5}, {-0.26, -0.27, 0.17, 0.87}}
	// bias := []float64{2, 3, .5}

	c, err := numpy.Dot(weights, inputs)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(c)

}
