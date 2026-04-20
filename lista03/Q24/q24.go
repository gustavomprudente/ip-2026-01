package main

import f "fmt"

func serie(A float64)float64 {
	return A - (A*A*A/6) + (A*A*A*A*A/120) - (A*A*A*A*A*A*A/5040)
}

func main () {
	var r, i float64 

	f.Printf("------TABELA------\nÂngulos   Valores\n")
	for i = 0; i <= 6.3; i += 0.1 {
		r = serie(i)
		f.Printf("  %.1f  ->  %.2f\n", i, r)
	}
}