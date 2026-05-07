package main

import "fmt"

func primo(A int) bool {
	if A < 2 {
		return false
	}

	for d := 2; d*d <= A; d++ {
		if A % d == 0 {
			 return false
		}
	}

	return true
}

func main () {
	var Primo bool

	n := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %d° número: ", i+1)
		fmt.Scan(&n[i])
	}	

	for i := 0; i < 10; i++ {
		Primo = primo(n[i]) 
		if Primo {
			fmt.Printf("%d é primo e está na %d posição do vetor.\n", n[i], i)
		}
	}
}