package main

import "fmt"

func main () {

	var m, s, contador int

	for i := 50; i <= 70; i += 2 {
		s += i
		contador++
	}
	m = s / contador
	fmt.Printf("Soma: %d\n", s)
	fmt.Printf("Média: %d\n", m)
}