package main 

import "fmt"

func main () {

	var i int

	fmt.Println("Digite a idade a ser analisada:")
	fmt.Scan(&i)

	switch {
	case i >= 0 && i <= 2:
		fmt.Printf("Para %d, a classificação é: Recém-nascido\n", i)
	case i >= 3 && i <= 11:
		fmt.Printf("Para %d, a classificação é: Criança\n", i)
	case i >= 12 && i <= 19:
		fmt.Printf("Para %d, a classificação é: Adolescente\n", i)
	case i >= 20 && i <= 55:
		fmt.Printf("Para %d, a classificação é: Adulto\n", i)
	case i >= 55:
		fmt.Printf("Para %d, a classificação é: Idoso\n", i)
	default:
		fmt.Println("Idade inválida\n")
	}
}
