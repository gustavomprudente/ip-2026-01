package main

import "fmt"

func main () {
	var (
		par, impar int
		impares []int
		pares []int
	)

	num := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %d° número: ", i+1)
		fmt.Scan(&num[i])
		if num[i]%2 == 0 {
			par += num[i]
			pares = append(pares, num[i])
		} else {
			impar++
			impares = append(impares, num[i])
		}
	}

	fmt.Printf("\nNúmeros pares digitados: %d\nSoma dos números pares digitados: %d\nNúmeros ímpares digitados: %d\nQuantidade de números ímpares digitados: %d\n", pares, par, impares, impar)
}