package main

import "fmt"

func main () {

	var fx, x float64

	fmt.Print("digite x: ")
	fmt.Scan(&x)

	fx = 8 / (2-x)
	fmt.Printf("O resultado é %.2f\n", fx)
}
