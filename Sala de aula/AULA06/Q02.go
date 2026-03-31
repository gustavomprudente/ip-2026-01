package main

import "fmt"

func main () {
	var v [5]float64
	var soma float64

	for i:= 0; i < 5; i++ {
		fmt.Printf("Digite o %d° número:", i+1)
		fmt.Scan(&v[i])
		soma += v[i]
	}
	
	fmt.Printf("A soma desses números é igual a: %.2f\n", soma)
}
