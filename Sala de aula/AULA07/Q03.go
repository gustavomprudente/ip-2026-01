package main

import f "fmt"

func main () {

	var x, y, z float64

	f.Println("Digite 3 números para calcular sua média:")
	f.Scan(&x, &y, &z)
	f.Printf("A média desses números é: %.2f\n", media(x, y, z))
}

func media (x, y, z float64) float64 {
	m := (x + y + z) / 3
	return m
}
