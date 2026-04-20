package main

import f "fmt"

func main () {
	var (
		N int
		R float64
	)

	f.Printf("Defina quantos termos serão gerados:\n")
	f.Scan(&N)
	if N <= 0 {
		f.Printf("Inválido.")
		return
	}

	for i := 1; i <= N; i++ {
		if i % 2 == 0 {
			R -= (1000 - ((float64(i-1)) * 3)) / float64(i)
		} else {
			R += (1000 - (float64(i-1)) * 3) / float64(i)
		}
	}
	f.Printf("Resultado = %.2f\n", R)
}