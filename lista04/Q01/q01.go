package main

import "fmt"

func pot(A, B int) int {

	if B == 0 {
		return 1
	}
	return A * pot(A, B-1)
}

func main () {
	var x, n, r int

	fmt.Scan(&x, &n)
	r = pot(x, n)
	fmt.Printf("x = %d\nn = %d\nx^n = %d\n", x, n, r)
}