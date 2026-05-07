package main

import "fmt"

func main () {

	var e1, e2, e3 int

	menor1, menor2, menor3 := 999999, 999999, 999999
	n := make([]int, 100)
	m := make([]int, 100)

	fmt.Println("\n       Insira os dados dos funcionários.\n(Digite dois números iguais a zero para terminar)\n")

	for i := 0; i < 100; i++ {
		fmt.Print("Número do empregado: ")
		fmt.Scan(&n[i])
		fmt.Print("Número de meses de trabalho deste empregado: ")
		fmt.Scan(&m[i])
		if m[i] == 0 && n[i] == 0 {
			fmt.Printf("Programa encerrado.\n")
			break
		}

		switch {
		case m[i] < menor1:
			menor3, e3 = menor2, e2
			menor2, e2 = menor1, e1
			menor1, e1 = m[i], n[i]

		case m[i] < menor2:
		menor3, e3 = menor2, e2
		menor2, e2 = m[i], n[i]

		case m[i] < menor3:
			menor3, e3 = m[i], n[i]
		}
	}
	fmt.Printf("Empregado mais recente: %d\nSegundo empregado mais recente: %d\nTerceiro empregado mais recente: %d", e1, e2, e3)
}