package main 

import (
	"fmt"
	"math"
)

func main () {
	var N, S float64
	b := true
	for i := 1; i <= 51; i += 2 {
		if b {
			S += 1 / float64(i*i*i)
		} else {
			S -= 1 / float64(i*i*i)
		}
		b = !b
	}
	N = math.Pow(S*32, 1.0/3)
	fmt.Printf("%.2f\n", N)
}