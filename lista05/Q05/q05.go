package main

import "fmt"

func main () {
	var X, P int
	num := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %d° número: ", i+1)
		fmt.Scan(&num[i])
		if i == 0{
			X = num[0]
		} else {
			if num[i] < X {
				X = num[i]
				P = i
			}
		}
	}
	fmt.Printf("\nO menor elemento do vetor é %d e sua posição dentro do vetor é: %d\n", X, P+!)
}