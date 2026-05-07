package main

import "fmt"

func main() {

	var r []int
	n := make([]int, 30)

	for i := 0; i < 30; i++ {
		fmt.Printf("Digite o %d° número: ", i+1)
		fmt.Scan(&n[i])

		switch {
		case i%2 == 0:
			r = append(r, n[i]*2)
		default:
			r = append(r, n[i]*3)
		}
	}

fmt.Printf("%d\n", r)
}