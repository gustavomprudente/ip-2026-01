package main

import "fmt"

func main() {

	var n1, n2, s, r int

	fmt.Println("Digite dois números:")
	fmt.Scan(&n1, &n2)
	s = n1 + n2

	switch {
		case s > 20:
			r = s + 8
			fmt.Printf("o resultado é %d\n", r)

		default:
			r = s - 5
			fmt.Printf("O resultado é %d\n", r)
	}
}
