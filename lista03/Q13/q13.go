package main

import "fmt"

func main () {
	var H, f float64

	f = 1
	for i := 1; i < 51; i++ {
		H += f / float64(i)
		f += 2
	}	
	fmt.Printf("H = %.2f\n", H)
}