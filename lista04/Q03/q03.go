package main

import "fmt"

func inverta (A []int, prim, ult int) {
	if prim >= ult {
		return
	}

	A[prim], A[ult] = A[ult], A[prim]
	inverta(A, prim+1, ult-1)
}

func main () {
	var N int

	fmt.Printf("Quantos números você deseja adicionar ao vetor?\n")
	fmt.Scan(&N)
	fmt.Printf("Digite os valores abaixo:\n")

	n := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&n[i])
	}

	inverta(n, 0, len(n)-1)
	fmt.Println(n)
}