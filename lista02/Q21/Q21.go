package main

import "fmt"

func main () {

	var (
		n int
		n1, n2, n3, m, a float64
		c, r string
	)

	fmt.Printf("número de identificação do aluno: ")
	fmt.Scan(&n)
	fmt.Printf("Nota 1: ")
	fmt.Scan(&n1)
	fmt.Printf("Nota 2: ")
	fmt.Scan(&n2)
	fmt.Printf("Nota 3: ")
	fmt.Scan(&n3)
	m = (n1 + n2 + n3) / 3
	a = (n1 + (2 * n2) + (n3 * 3) + m) / 7
	switch {
	case a >= 9.0 && a <= 10.0:
		c = "A"
		r = "Aprovado"
	case a >= 7.5 && a < 9.0:
		c = "B"
		r = "Aprovado"
	case a >= 6.0 && a < 7.5:
		c = "C"
		r = "Aprovado"
	case a >= 4.0 && a < 6.0:
		c = "D"
		r = "Reprovado"
	case a < 4.0:
		c = "E"
		r = "Reprovado"
	default:
		fmt.Println("Nota inválida")
	}
	fmt.Print("-------------------------------------")
	fmt.Printf("número de identificação do aluno: %d\nnota 1: %.2f\nnota 2: %.2f\nnota 3: %.2f\nMédia dos exercícios: %.2f\nMédia de aproveitamento: %.2f\nConceito correspondente: %s\n%s\n", n, n1, n2, n3, m, a, c, r)
}
