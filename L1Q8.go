package main

import (
	"fmt"
	"math"
)

func main() {
	var r, h, c, at, ab, al float64
	fmt.Scan(&r)
	fmt.Scan(&h)
	ab = math.Pi * math.Pow(r, 2)
	al = 2 * math.Pi * r * h
	at = (2 * ab) + al
	c = at * 100
	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", c)
}