package main

import "fmt"

func soma(A []float64, B int) float64 {
	if B == 0 {
		return A[0]
	}
	
	return A[B] + soma(A, B-1)
}

func main () {
	var N int

	fmt.Printf("Quantos números você deseja somar?\n")
	fmt.Scan(&N)
	fmt.Printf("Digite os valores abaixo:\n")

	n := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&n[i])
	}
	fmt.Printf("Soma dos valores: %.2f", soma(n, len(n)-1))
}