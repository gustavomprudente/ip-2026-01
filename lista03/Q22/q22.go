package main

import "fmt"

func main () {
	var  S float64

	for i := 1; i <= 37; i++ {
		S += float64(38 - i) * float64(39 -i) / float64(i)
	}
	fmt.Printf("%.2f\n", S)
}