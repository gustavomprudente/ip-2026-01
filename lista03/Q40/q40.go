package main

import "fmt"

func main() {
	var L, ing, Lmax, ingmax, pmax float64
	fmt.Printf("-----------------------------TABELA------------------------------\n")

	ing = 130
	for i := 6.0; i > 1; i -= 0.6 {
		L = (i * ing) - 300
		fmt.Printf("Para %.0f ingressos, cada um a %.1f reais, temos um lucro de %.2f\n", ing, i, L)
		ing += 30
		if i == 6 {
			pmax = i
			Lmax = L
			ingmax = ing
		}
		if L > Lmax {
			pmax = i
			Lmax = L
			ingmax = ing
		}
	}
	fmt.Printf("\nO lucro máximo será de %.2f, para %.0f ingressos cada um a %.1f reais.\n", Lmax, ingmax, pmax)
}