package main

import "fmt"

func main() {
	var soma float32
	alt := make([]float32, 10)

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %d° número: ", i+1)
		fmt.Scan(&alt[i])
		soma += alt[i]
	}

	m := soma / 10 

	fmt.Println()
	for i := 0; i < 10; i++ {
		if alt[i] > m {
			fmt.Printf("%.2f ", alt[i])
		}
	}
}