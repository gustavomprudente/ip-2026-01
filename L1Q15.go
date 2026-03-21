package main

import(
	"fmt"
)

func main () {
	var n, x, r int
	fmt.Scan(&n)
	for x = 2; x <= n; x += 2 {
		r = x * x
		fmt.Printf("%d^2 = %d\n", x, r)
	}
}