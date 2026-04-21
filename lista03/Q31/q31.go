package main

import (
	"fmt"
	"math"
)

func main () {
	var grãos float64
	for i := 0.0; i < 64; i++ {
		grãopquad := math.Pow(2, i)
		grãos += grãopquad
	}

	fmt.Printf("%.2f", grãos)
}