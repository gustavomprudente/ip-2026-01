package main

import "fmt"

func main () {
	var (
		N, a, e, r int
		n1, n2, m, soma float64
	)

	fmt.Printf("Quantos alunos serão analisados?\n")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("--Aluno %d--\n", i+1)
		fmt.Printf("Nota 1: ")
		fmt.Scan(&n1)
		fmt.Printf("Nota 2: ")
		fmt.Scan(&n2)

		m = (n1 + n2) / 2
		fmt.Printf("Média do aluno %d: %.2f\n", i+1, m)

		switch {
		case m <= 3:
			fmt.Printf("Reprovado.\n")
			r++
		case m > 3 && m < 7:
			fmt.Printf("Exame.\n")
			e++
		case m >= 7:
			fmt.Printf("Aprovado.\n")
			a++			
		}
		soma += n1+n2
	}

	média := soma / float64(2*N)
	fmt.Printf("\nTotal de alunos aprovados: %d\n", a)
	fmt.Printf("Total de alunos de exame: %d\n", e)
	fmt.Printf("Total de alunos reprovados: %d\n", r)
	fmt.Printf("Média da classe: %.2f\n", média)
}