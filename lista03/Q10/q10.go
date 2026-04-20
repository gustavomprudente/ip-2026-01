package main

import "fmt"

func main () {
	var N int

	fmt.Printf("Quantos termos terá a sequência?\n")
	fmt.Scan(&N)
	f := make([]int, N)
	f[0] = 0
	f[1] = 1
	for i := 0; i < N - 2; i++ {
		f[i+2] = f[i] + f[i+1]
	}
	for i := 0; i < N; i++ {
		if i == N-1 {
			fmt.Printf("%v\n", f[i])
		} else {
			fmt.Printf("%v - ", f[i])
		}
	}	
}