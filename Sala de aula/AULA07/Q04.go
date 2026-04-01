package main 

import f "fmt"

func main () {
	var n int

	f.Println("Digite um número para calcular seu fatorial.")
	f.Scan(&n)
	f.Printf("O fatorial de %d é %d", n, fatorial(n))
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	} 
	return n * fatorial(n-1)
}
