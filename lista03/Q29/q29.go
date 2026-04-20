package main

import f "fmt"

func main() {
	var N, S int

	f.Printf("Digite N: ")
	f.Scan(&N)
	for i := 1; i <= N; i++ {
		S += i
	}
	f.Printf("O somatório de 1 até %d = %d\n", N, S)
}