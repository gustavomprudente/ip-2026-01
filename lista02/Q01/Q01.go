package main 

import "fmt"

func main () {

	var n int

	fmt.Println("Digite um número para verificar se ele é par ou ímpar:")
	fmt.Scan(&n)
	if n % 2 == 0 {
		fmt.Printf("O número informado é par\n")
	} else {
		fmt.Printf("O número informado é ímpar\n")
	}
}
