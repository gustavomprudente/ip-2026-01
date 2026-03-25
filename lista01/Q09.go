package main

import(
	"fmt"
	"math"
)

func main () {
	var a, b, c, delta float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	delta = math.Pow(b,2) - 4 * a * c
	fmt.Printf("O VALOR DE DELTA E = %.2f\n", delta)

}
