package main

import "fmt"

func main () {
	var n1, n2, resto int
	quociente := 0
	fmt.Printf("Digite n1 e n2:\n")
	fmt.Scan(&n1, &n2)
	if n1 < 0 || n2 < 0 {
		fmt.Printf("Dígito inválido.")
		return
	}
	if n2 > n1 {
		n1, n2 = n2, n1
	}

	for resto = n1; ; resto -= n2 {
		if resto < n2 {
			break
		}
		quociente++
	}
	fmt.Printf("Quociente(%d, %d) = %d\nResto(%d, %d) = %d", n1, n2, quociente, n1, n2, resto)
}