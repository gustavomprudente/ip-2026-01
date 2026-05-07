package main

import (
	"fmt"
	"math"
)

func main () {

	var S float64
	n := make([]float64, 100)

	for i := 0; i < 100; i++ {
		fmt.Printf("Digite o %d° número: ", i+1)
		fmt.Scan(&n[i])
	}	

	for i := 0; i < 50; i++ {
		S += math.Pow((n[i] - n[99-i]), 3)
	}

	fmt.Printf("%.2f", S)
}