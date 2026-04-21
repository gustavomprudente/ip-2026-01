package main

import "fmt"

func main () {
	var n int

	fmt.Printf("Digite um número para transformá-lo em base 16: ")
	fmt.Scan(&n)
	fmt.Printf("O número %d na base 16 é: %x\n", n, n)
}