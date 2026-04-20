package main

import "fmt"

func fat(A int)int {
	if A == 0 {
		return 1
	}
	return A * fat(A-1)
}

func somatório(X float64) float64 {
	var S float64
	for i := 2; i < 20; i++ {
		if i % 2 == 0 {
			S += (X / float64(fat(i)))
		} else {
			S -= (X / float64(fat(i)))
		}
	}
	return S
}

func main () {
	var x, resultado float64

	fmt.Printf("x = ")
	fmt.Scan(&x)
	resultado = somatório(x)
	fmt.Printf("%.2f\n", resultado)
}