package main

import "fmt"

func main () {
	var tc, mm, tf, pol float64
	fmt.Scan(&tf)
	fmt.Scan(&pol)
	tc = 5 * (tf - 32) / 9
	mm = pol * 25.4
	fmt.Printf("O VALOR EM CELSIUS = %.2f\nA QUANTIDADE DE CHUVA E = %.2f", tc, mm)
}
