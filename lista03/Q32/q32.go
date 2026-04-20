package main

import "fmt"

func main() {
	var n1, n2, r int

	fmt.Print("N1: ")
	fmt.Scan(&n1)
	fmt.Print("N2: ")
	fmt.Scan(&n2)

	negativo := false
	if n1 < 0 {
		n1 = -n1
		negativo = !negativo
	}
	if n2 < 0 {
		n2 = -n2
		negativo = !negativo
	}

	for i := 0; i < n2; i++ {
		r += n1
	}

	if negativo {
		r = -r
	}

	fmt.Printf("Resultado: %d\n", r)
}