package main

import "fmt"

func main () {
	var peso, gordo, magro float64
	var n, nmaior, nmenor, escolha int
	vezes := 90
	fmt.Printf("Deseja digitar todos os 90 bois?\n1 para sim\n2 para não\n")
	fmt.Scan(&escolha)
	if escolha == 2 {
		fmt.Printf("Escolha quantos bois você deseja adicionar: ")
		fmt.Scan(&vezes)
	}

	for i := 0; i < vezes; i++ {
		fmt.Printf("--boi %d--\nn° identificador: ", i+1)
		fmt.Scan(&n)
		fmt.Printf("Peso: ")
		fmt.Scan(&peso)
		switch {
		case i == 0:
    		gordo = peso
			magro = peso
			nmaior = n
			nmenor = n
		case peso > gordo:
			gordo = peso
			nmaior = n
		case peso < magro:
			magro = peso
			nmenor = n
		}
	}
	fmt.Printf("\nBoi mais gordo:\nIdentificação: %d\nPeso: %.2f\n", nmaior, gordo)
	fmt.Printf("Boi mais magro:\nIdentificação: %d\nPeso: %.2f\n", nmenor, magro)
}