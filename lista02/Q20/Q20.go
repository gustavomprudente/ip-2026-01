package main

import "fmt"

func main () {

	var (
		mp int
		v, vf float64
	)
	fmt.Printf("Qual o preço normal de etiqueta do produto?\n")
	fmt.Scan(&v)
	fmt.Printf("Escolha a condição de pagamento:\n1 - À vista, dinheiro ou cheque, 10%% de desconto\n2 - À vista, cartão de crédito, 5%% de desconto\n3 - Em 2 vezes, preço normal de etiqueta sem juros\n4 - Em 3 vezes, preço normal de etiqueta + 10%% de jurosn\n")
	fmt.Scan(&mp)

	switch mp {
	case 1:
		vf = 0.9 * v
		fmt.Printf("Valor a ser pago: R$ %.2f\n", vf)
	case 2:
		vf = 0.95 * v	
		fmt.Printf("Valor a ser pago: R$ %.2f\n", vf)
	case 3:
		vf = 0.5 * v
		fmt.Printf("Valor a ser pago: 2 parcelas de R$ %.2f\n", vf)
	case 4:
		vf = (v * 1.1) / 3
		fmt.Printf("Valor a ser pago: 3 parcelas de R$ %.2f\n", vf)
	default:
		fmt.Printf("Valor inválido\n")	
	}
}
