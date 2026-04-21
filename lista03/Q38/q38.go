package main

import "fmt"

func main () {
	var soma, dverif1, dverif2 int
	cpf := make([]int, 11)
	fmt.Printf("Digite o cpf para verificá-lo. (dê espaço entre os dígitos)\n")
	for i := 0; i <= 10; i++ {
		fmt.Scan(&cpf[i])
	}

	for i := 0; i < 9; i++ {
		soma += cpf[i] * (10-i)
	}
	if (soma%11) < 2 {
		dverif1 = 0
	} else {
		dverif1 = 11 - (soma%11)
	}
	if dverif1 != cpf[9] {
		fmt.Printf("O dígito verificador 1 não é válido. O correto é: %d\n", dverif1)
		cpf[9] = dverif1
	}

	soma = 0
	for i := 0; i < 10; i++ {
		soma += cpf[i] * (11-i)
	}
	dverif2 = 11 - (soma % 11)

	if (dverif1 == cpf[9]) && (dverif2 == cpf[10]) {
		fmt.Printf("O cpf é válido.\n")
	} else {
		cpf[10] = dverif2
		fmt.Printf("O cpf é inválido. O correto seria: ")
		for i := 0; i <= 10; i++ {
			fmt.Printf("%d", cpf[i])
		}
	}
}