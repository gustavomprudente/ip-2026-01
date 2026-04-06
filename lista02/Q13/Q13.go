package main

import "fmt"

func main () {

	var n1, n2 int

	fmt.Println("Digite um número: ")
	fmt.Scan(&n1)

	if n1 >= 100 && n1 <= 999 {
		n2 = (n1 / 10) % 10
		fmt.Printf("o algarismo da casa das dezenas é %d\n", n2)
	} else {
		fmt.Println("Número inválido. O número precisa ter 3 casas.\n")
	}
}
