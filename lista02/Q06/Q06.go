package main

import "fmt"

func main () {

	var x int

	fmt.Println("Digite um número para verificar se ele é divisível por 5")
	fmt.Scan(&x)

	if x % 5 == 0 {
		fmt.Printf("O número digitado é divisível por 5\n")
	} else {
		fmt.Printf("O número digitado não é divisível por 5\n")
	}
}
