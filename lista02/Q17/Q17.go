package main

import "fmt"

func main() {

	var (
		consumo, v float64
		conta int
		tc string
	)

	fmt.Print("Conta: ")
	fmt.Scan(&conta)
	fmt.Print("Tipo da conta: ")	
	fmt.Scan(&tc)
	fmt.Print("Consumo (m³): ")
	fmt.Scan(&consumo)

	switch {
	case (tc == "C")||(tc == "Comercial")||(tc == "comercial")||(tc == "c"):
		if consumo >= 80 {
			v = 500 + ((consumo - 80) * 0.25)
		} else {
			v = 500
		}
	case (tc == "I")||(tc == "Industrial")||(tc == "industrial")||(tc == "i"):
		if consumo >= 100 {
			v = 800 + ((consumo - 100) * 0.04)
		} else {
			v = 800
		}
	case (tc == "R")||(tc == "Residencial")||(tc == "residencial")||(tc == "r"): 
		v = 5 + (0.05 * consumo)
	default:
		fmt.Println("erro\n")
		return
	}
	fmt.Println("----------------------------")
	fmt.Printf("CONTA = %d\nVALOR DA CONTA = %.2f\n", conta, v)
}
