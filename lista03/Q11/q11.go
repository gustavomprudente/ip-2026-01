package main

import "fmt"

func fat(A int)int {
	if A == 0 {
		return 1
	}
	return A * (A-1)
}
func main () {
	var n, r int

	fmt.Printf("Digite um número para calcular seu fatorial.\n")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Printf("Número inválido\n")
		return
	}
	r = fat(n)
	fmt.Printf("%d! = %d\n", n, r)
}