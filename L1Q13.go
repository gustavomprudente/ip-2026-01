package main

import"fmt"

func main() {
	var x float32
	fmt.Scan(&x)
	if (x >= 0) && (x < 6) {
		fmt.Printf("NOTA = %2.f CONCEITO = D\n", x)
	} else if (x >= 6) && (x < 7.5) {
		fmt.Printf("NOTA = %2.f CONCEITO = C\n", x)
	} else if (x >= 7.5) && (x < 9) {
		fmt.Printf("NOTA = %2.f CONCEITO = B\n", x)
	} else if (x >= 9) && (x <= 10) {
		fmt.Printf("NOTA = %2.f CONCEITO = A\n", x)
	} else {
		fmt.Print("erro\n")
	}
}