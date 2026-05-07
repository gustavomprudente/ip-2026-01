package main

import "fmt"

func main () {

	f := make([]int, 50)
	f[0] = 0
	f[1] = 1
	for i := 0; i < 50 - 2; i++ {
		f[i+2] = f[i] + f[i+1]
	}
	fmt.Printf("%d", f)
}