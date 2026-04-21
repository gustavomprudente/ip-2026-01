package main

import "fmt"

func main() {
	var n1, n2, mmc int

	fmt.Printf("Digite n1 e n2:\n")
	fmt.Scan(&n1, &n2)
	mmc = 1
	for mmc%n1 != 0 || mmc%n2 != 0 {
		mmc++
	}

	fmt.Printf("mmc(%d e %d) = %d\n", n1, n2, mmc)	
}