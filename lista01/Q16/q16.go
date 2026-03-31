package main

import "fmt"

func main () {
	var R, S float64
	fmt.Println("DIGITE O SALÁRIO PARA SER REAJUSTADO:")
	fmt.Scan(&S)
	if S > 300 {
		R = S * 1.3
		fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", R)
	} else if (S <= 300) && (S >= 0) {
		R = S * 1.5
		fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", R)
	} else {
		fmt.Println("erro\n")
	}
}
