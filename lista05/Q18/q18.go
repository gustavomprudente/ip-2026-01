package main

import "fmt"

func Sort (A []int) []int {
	for i := 0; i < len(A); i++ {
		menor := i
		for j := i+1; j < len(A); j++ {
			if A[j] < A[menor] {
				menor = j
			}
		}
		A[i], A[menor] = A[menor], A[i]
	}
	return A
}

func main () {

	n := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d° número: ", i+1)
		fmt.Scan(&n[i])
	}

	n = Sort(n)
	fmt.Println(n)
}