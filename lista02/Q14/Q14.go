package main

import "fmt"

func main () {

	var (
		a, b, c, d, pi, pf float64
		n1, n2, n3, n4 string
	)

	fmt.Printf("Informe o preço inicial de fábrica: R$ ")
	fmt.Scan(&pi)
	
	fmt.Printf("Agora, escolha as opções para o carro (s ou n):\na) Ar condicionado (R$ 1750,00)\n")
	fmt.Scan(&n1)
	fmt.Printf("b) Pintura metálica (R$ 800,00)\n")
	fmt.Scan(&n2)
	fmt.Printf("c) Vidro elétrico (R$ 1200,00)\n")
	fmt.Scan(&n3)
	fmt.Printf("d) Direção hidráulica (R$ 2000,00)\n")
	fmt.Scan(&n4)
	if n1 == "s" {
		a = 1750
	} else {
		a = 0
	}

	if n2 == "s" {
		b = 800
	} else {
		b = 0
	}

	if n3 == "s" {
		c = 1200
	} else {
		c = 0
	}

	if n4 == "s" {
		d = 2000
	} else {
		d = 0
	}
	pf = pi + a + b + c + d
	fmt.Printf("Preço final do carro: R$ %.2f\n", pf)	
}
