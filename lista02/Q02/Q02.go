package main 

import "fmt"

func main () {

	var n int

	fmt.Println("Digite um número para verificar se ele é positivo, negativo ou nulo:")
	fmt.Scan(&n)
	
	switch {
	case n > 0:
		fmt.Println("O valor é positivo\n")
	case n == 0:
		fmt.Println("O valor é nulo\n")
	case n < 0:
		fmt.Println("O valor é negativo\n")	
	}
}
