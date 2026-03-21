package main 

import "fmt"

func main () {
	var a1, r, n, s int
	fmt.Scan(&a1, &r, &n)
	s = (2 * a1 + (n - 1) * r) * n / 2
	fmt.Printf("%d\n", s)
}