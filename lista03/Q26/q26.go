package main

import f "fmt"

func fat(A int)int {
	if A == 0 {
		return 1
	}
	return A * fat(A-1)
}

func main () {
	var S float64
	for i := 0; i <= 19; i++ {
		S += (100-float64(i)) / float64(fat(i))
	}
	f.Printf("S = %.2f\n", S)
}