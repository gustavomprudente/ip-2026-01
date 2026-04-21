package main

import (
	"fmt"
	"math"
)

func main() {
	var n, r int
	seq := []int{}
	fmt.Printf("Digite um número na base 8 para transformá-lo em um número na base 10\n")
	fmt.Scan(&n)

	for n > 0 {
		digito := n % 10
		seq = append(seq, digito)
		n = n / 10
	}

	for i := range seq {
		r += int(math.Pow(float64(8), float64(i))) * seq[i]
	}

	fmt.Printf("O número %d na base 10 é: %d", n, r)
}