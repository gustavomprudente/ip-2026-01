package main

import "fmt"

func main() {

	var A, B, C, MAIOR, MENOR, INTER int

	fmt.Println("Digite três valores para colocá-los em ordem crescente.")
	fmt.Scan(&A, &B, &C)

	switch {
	case (A > B && A > C):
		MAIOR = A
		if B > C {
			INTER = B
			MENOR = C
		} else {
			INTER = C
			MENOR = B
		}

	case (B > A && B > C):
		MAIOR = B
		if A > C {
			INTER = A
			MENOR = C
		} else {
			INTER = C
			MENOR = A
		}

	case (C > A && C > B):
		MAIOR = C
		if A > B {
			INTER = A
			MENOR = B
		} else {
			INTER = B
			MENOR = A
		}

	default:
		fmt.Println("Há valores iguais ou erro\n")
		return
	}

	fmt.Printf("Ordem crescente: %d %d %d\n", MENOR, INTER, MAIOR)
}
