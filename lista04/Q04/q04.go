package main

import "fmt"

func binario(A int) {
	if A == 0 {
		return
	}

	binario(A / 2)
	fmt.Print(A % 2)
}

func main() {
	var n int

	fmt.Print("Digite um número decimal: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Print("0")
	} else {
		binario(n)
	}
}