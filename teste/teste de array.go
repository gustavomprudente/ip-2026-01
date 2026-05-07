package main

import (
	"fmt"
	"sort"
)

func buscaBinaria(N int, p int) int {
	for i := 0; i < N; i++ {
		switch {
		case N[(len(N)-1)/2] > p:
		
		case N[(len(N)-1)/2] < p:
	
		case N == p {
			return i
		}
	}
}

func main () {
	var n [5]float64
	var busca int

	for i := 0; i < 5; i++ {
		fmt.Printf("n° %d", i+1)
		fmt.Scan(&n[i])
		if n == 4 {
			fmt.Printf("Qual número você deseja procurar?")
			fmt.Scan(&procura)
		}
		sort.Ints(n)
		fmt.Printf("%d\n\n", n)
	}
	busca = buscaBinaria(n, procura)
}