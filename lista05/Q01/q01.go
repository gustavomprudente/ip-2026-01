package main

import "fmt"

func main () {
	var (
		sup []int
		pos int
		posSup []int
	)

	num := make([]int, 10)
	
	for pos = 0; pos < 10; pos++ {
		fmt.Printf("Digite o %d° número: ", pos+1)
		fmt.Scan(&num[pos])
		fmt.Println()
		if num[pos] > 50 {
			sup = append(sup, num[pos])
			posSup = append(posSup, pos)
		}
	}
	
	if len(sup) == 0 {
		fmt.Printf("Não há nenhum número superior a 50.")
		return
	}

	for i := 0; i < len(sup); i++ {
		fmt.Printf("O número %d é maior que 50 e está na posição %d do vetor.\n", sup[i], posSup[i])
	}
}