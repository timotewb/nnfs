package app

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

func GaussianRandom(mean float64, variance float64) float64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	u1 := rand.Float64()
	u2 := rand.Float64()
	r := math.Sqrt(-2 * math.Log(u1))
	theta := 2 * math.Pi * u2
	z := r * math.Cos(theta)

	// Adjust the generated number to have the desired mean and variance.
	adjustedZ := z*math.Sqrt(variance) + mean
	return adjustedZ
}

func GenerateGaussian() {
	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < 10000; i++ {
		value := GaussianRandom(0, 1)
		_, err := file.WriteString(fmt.Sprintf("%f\n", value))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
