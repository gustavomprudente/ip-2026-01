package main

import "fmt"

func main () {
	var n int

	raiz := make([]int, 15)

	for i := 0; i < 15; i++ {
		fmt.Printf("Digite o %d° número: ", i+1)
		fmt.Scan(&n)
		if n < 0 {
			raiz[i] = -1
		} else {
			raiz[i] = n * n
		}
	}	

	fmt.Printf("Valores armazenados: %d\n", raiz)
}