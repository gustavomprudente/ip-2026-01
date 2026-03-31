package main

import "fmt"

func main() {
	var conta int
	var consumo, v float32
	var tc string

	fmt.Scan(&conta, &consumo, &tc)

	if tc == "C" {
		if consumo >= 80 {
			v = 500 + ((consumo - 80) * 0.25)
		} else {
			v = 500
		}
	} else if tc == "I" {
		if consumo >= 100 {
			v = 800 + ((consumo - 100) * 0.04)
		} else {
			v = 800
		}
	} else if tc == "R" {
		v = 5 + (0.05 * consumo)
	} else {
		fmt.Println("erro")
	}

	fmt.Printf("CONTA = %d\nVALOR DA CONTA = %.2f\n", conta, v)
}
