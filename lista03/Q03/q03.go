package main

import "fmt"

func main () {

	var salarioCarlos, salarioJoao float64
	m := 0

	fmt.Printf("Salário de Carlos: ")
	fmt.Scan(&salarioCarlos)
	salarioJoao = salarioCarlos / 3

	for i := 0; salarioCarlos >= salarioJoao; i++ {
		salarioCarlos = salarioCarlos * 1.02
		salarioJoao = salarioJoao * 1.05
		m++
	}
	fmt.Printf("O salário de João ultrapassará ou igualará o de Carlos em %d meses. %d\n", m)
	fmt.Printf("Salário João: %.2f\n", salarioCarlos)
	fmt.Printf("Salário João: %.2f\n", salarioJoao)
}