package main

import "fmt"

func inverso (A []float64, prim, ult int) {
	if prim >= ult {
		return
	}

	A[prim], A[ult] = A[ult], A[prim]
	inverso(A, prim+1, ult-1)
}

func main () {
	
	c := 1
	fmt.Printf("Código: ")
	fmt.Scan(&c)
	if c == 0 {
		fmt.Println("Programa encerrado.")
		return
	}

	fmt.Printf("Digite os valores do vetor:\n")
	n := make([]float64, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&n[i])
	}

	if c == 1 {
		fmt.Println(n)
	}

	if c == 2 {
		inverso(n, 0, 9)
		fmt.Println(n)
	}
}