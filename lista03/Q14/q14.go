package main

import "fmt"

func primo(A int) bool {
	for d := 2; d < A; d++ {
		if A % d == 0 {
			 return false
		}

	}
	return true
}

func main () {
	var (
		n1, n2 int
		r []int
	)
	fmt.Printf("Digite o intervalo que será avaliado.\n")
	fmt.Scan(&n1, &n2)
	
	if n1 < 0 || n2 < 0 {
		fmt.Printf("Dígito inválido.\n")
		return
	}

	for i := n1; i <= n2; i++ {
		primo := primo(i)
		if primo {
			r = append(r, i)
		}
	}
	fmt.Printf("%v", r)
}