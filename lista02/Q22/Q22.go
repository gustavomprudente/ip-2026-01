package main

import "fmt"

func main () {

	var (
		mat, he int
		sm float64 = 788.0
		vhe float64 = 10.0

		se, sb, inss, ir, sl float64
	)

	fmt.Printf("Matrícula do funcionário: ")
	fmt.Scan(&mat)
	fmt.Printf("Quantidade de horas-extras trabalhadas: ")
	fmt.Scan(&he)
	se = float64(he) * vhe
	sb = (3 * sm) + se

	if sb > 1500 {
		inss = 0.12 * sb
	}

	if sb > 2000 {
		ir = 0.20 * sb
	}
	sl = sb - inss - ir
	
	fmt.Printf("---------------------\nMatrícula: %d\nSalário Bruto: R$ %.2f\nDesconto INSS: R$ %.2f\nDesconto IR: R$ %.2f\nSalário Líquido: R$ %.2f\n", mat, sb, inss, ir, sl)
}
