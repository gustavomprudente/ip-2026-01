package main

import "fmt"

func main () {

	var d, i, c float32

	fmt.Println("Insira o destino:\n1 - Região Norte\n2 - Região Nordeste\n3 - Região Centro-Oeste\n4 - Região Sul\n")
	fmt.Scan(&d)
	fmt.Println("A viagem inclui Retorno?\n1 - sim\n2 - não\n")
	fmt.Scan(&i)
	switch d {
	case 1:
		if i == 1 {
			c = 900
		} else {
			c = 500
		}
	case 2:
		if i == 1 {
			c = 650
		} else {
			c = 350
		}
	case 3:
		if i == 1 {
			c = 600
		} else {
			c = 350
		}
	case 4:
		if i == 1 {
			c = 550
		} else {
			c = 300
		}
	default:
		fmt.Println("Dígito inválido.\n")
		return
	}
	fmt.Printf("Preço da passagem: R$ %.2f\n", c)
}
