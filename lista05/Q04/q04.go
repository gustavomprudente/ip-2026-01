package main

import "fmt"

func main () {
	var (
		rep, contador []int
		repete int
	)
	num := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %d° número: ", i+1)
		fmt.Scan(&num[i])
	}
	
	for i := 0; i < len(num); i++ {
		cont := 0
		for p := i+1; p < len(num); p++ {
			if num[i] == num[p] {
				cont++
				num = append(num[:p], num[p+1:]...)
				p--
				repete = num[i]
			}
		}
		if cont > 0 {
			contador = append(contador, cont)
			rep = append(rep, repete)
		}
	}

	for i := 0; i < len(contador); i++ {
		fmt.Printf("\nO número %d está repetido e se repete %d vezes\n", rep[i], contador[i])
	}
}