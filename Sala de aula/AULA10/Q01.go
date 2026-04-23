package main

import "fmt"

type Pessoa struct {
	nome      string
	altura    float64
	pesoideal float64
}

func main() {
	var (
		pes    []Pessoa
		n      string
		a      float64
		pideal float64
	)

	for {
		fmt.Printf("(para terminar o programa digite 'FIM')\nNome da pessoa: ")
		fmt.Scan(&n)
		if n == "FIM" {
			break
		}
		fmt.Printf("Altura da pessoa: ")
		fmt.Scan(&a)
		fmt.Println()
		pideal = 72.7*a - 58.0
		pes = append(pes, Pessoa{
			nome:      n,
			altura:    a,
			pesoideal: pideal,
		})
	}
	fmt.Println("---Dados---")

	for i, p := range pes {
		fmt.Printf("Pessoa %d\n", i+1)
		fmt.Printf("Nome: %s\n", p.nome)
		fmt.Printf("Altura: %.2f\n", p.altura)
		fmt.Printf("Peso ideal: %.2f\n\n", p.pesoideal)
	}
}