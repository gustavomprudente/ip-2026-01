package main

import "fmt"

func main () {

	var i int

	fmt.Printf("Digite a idade da pessoa a ser analisada:\n")
	fmt.Scan(&i)

	switch {
	case i < 16:
		fmt.Println("Não-eleitor")
	case i >= 18 && i <= 65:
		fmt.Println("Eleitor obrigatório")
	case (i >= 16 && i < 18) || i > 65:
		fmt.Println("Eleitor facultativo")
	default:
		fmt.Println("Idade inválida")
	}
}
