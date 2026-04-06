package main 

import (
	"fmt"
	s "strings"
)

func main () {

	var (
		d int
		p, pf float32
		cat string
	)

	fmt.Printf("Preço normal do DVD alugado: ")
	fmt.Scan(&p)
	fmt.Printf("Categoria (comum ou lançamento): ")
	fmt.Scan(&cat)
	cat = s.ToLower(cat)
	fmt.Printf("Digite\n1 para Domingo\n2 para Segunda\n3 para Terça\n4 para Quarta\n5 para Quinta\n6 para Sexta\n7 para Sábado\nDia do aluguel:")
	fmt.Scan(&d)

	switch d {
	case 2, 3, 5:
		p = 0.6 * p
		switch cat {
		case "comum", "normal":
			pf = p
		case "lançamento": 
			pf = 1.15 * p
		}
	case 4, 6 ,7, 1:
		switch cat {
		case "comum", "normal":
			pf = p
		case "lançamento":
			pf = 1.15 * p
		}
	default:
		fmt.Printf("Inválido.\n")
		return
	}
	fmt.Printf("A locação do dvd será de R$ %.2f\n", pf)
}
