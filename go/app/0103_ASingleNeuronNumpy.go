package app

import (
	"fmt"
	"log"

	"github.com/timotewb/gonn/numpy"
)

func ASingleNeuronNumpy() {
	fmt.Println(" 01 ASingleNeuronNumpy()")

	inputs := []float64{1, 2, 3, 2.5}
	weights := []float64{.2, .8, -.5, 1}
	// bias := 2

	c, err := numpy.Dot(inputs, weights)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(c)
}
