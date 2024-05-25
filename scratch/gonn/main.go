package main

import (
	"fmt"

	"github.com/timotewb/gonn/numpy/random"
)

func main() {
	m := random.Randn(1, 2)
	fmt.Println(m)
}
