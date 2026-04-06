package main

import (
	"fmt"
	"math"
)

func main() {

	var n1, r float64

	fmt.Println("Digite um número para calcular sua raiz quadrada, caso seja positivo ou zero, ou seu quadrado, se for negativo.")
	fmt.Scan(&n1)

	if n1 >= 0 {
		r = math.Pow(n1, 0.5)
		fmt.Printf("O resultado é %f\n", r)

	} else {
		r = n1 * n1
		fmt.Printf("O resultado é %f\n", r)
	}
}
