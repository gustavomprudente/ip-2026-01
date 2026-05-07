package main

import "fmt"

func FA(A int, notas []int) int {
	contador := 0
	for i := 0; i < 15; i++ {
		if A == notas[i] {
			contador++
		}
	}
	return contador
}
func main () {
	
	var (
		fa []int
		fr []float64
	)

	
	nota := make([]int, 15)
	for i := 0; i < 15; i++ {
		fmt.Printf("%d° nota: ", i+1)
		fmt.Scan(&nota[i])
	}

	for i := 0; i < 15; i++ {
		fa = append(fa, FA(nota[i], nota))
		fr = append(fr, float64(fa[i]) / 15 * 100)
	}
	fmt.Println("-----------------------------TABELA-----------------------------\n   Nota         Frequência absoluta          Frequência relativa")
	for i := 0; i < 15; i++ {
		fmt.Printf("    %d                   %d                          %.2f%%\n", nota[i], fa[i], fr[i])
	}
}