package main

import "fmt"

func main() {
	var n, i int
	var s float64

	fmt.Println("Digite o valor de n: ")
	fmt.Scan(&n)

	if n > 0 {
		for i = 1; i <= n; i++ {
			s += 1 / float64(i)
		}
		fmt.Printf("%.6f\n", s)
	} else {
		fmt.Println("erro\n")
	}
}