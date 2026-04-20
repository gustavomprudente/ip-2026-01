package main

import (
	"math"
	"fmt"
)

func fat(A int)int {
	if A == 0 {
		return 1
	}
	return A * fat(A-1)
}

func main () {
	var x float64

	fmt.Printf("Valor de X: ")
	fmt.Scan(&x)

	b := true
	cosseno := 1.0 
	for i := 2; i <= 38; i += 2 {
		if b {
			cosseno -= math.Pow(x, float64(i)) / float64(fat(i))
		} else {
			cosseno += math.Pow(x, float64(i)) / float64(fat(i))
		}
		b = !b
	}
	dif := cosseno - math.Abs(math.Cos(x))

	fmt.Printf("Cosseno calculado: %f\n", cosseno)
	fmt.Printf("diferença: %f\n", dif)
}