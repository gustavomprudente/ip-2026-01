package main

import (
	"math"
	"fmt"
)

func main () {
	for i := 0.0; i <= 20; i += 0.5 {
		V := 4.0/3 * math.Pi * math.Pow(i, 3)
		fmt.Printf("Volume com raio %.1f cm = %.2f cm³\n", i, V)
	}
}