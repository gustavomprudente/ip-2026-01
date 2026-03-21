package main

import "fmt"

func main(){
	var v, h int
	fmt.Scan(&h)
	v = 10 * (h / 3) + ((h % 3) * 5)
	fmt.Printf("O VALOR A PAGAR E = %.2f\n", float32(v))
}