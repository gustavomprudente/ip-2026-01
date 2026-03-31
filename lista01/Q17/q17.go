package main

import "fmt"

func main () {
	var x, y, z int
	fmt.Scan(&x, &y)
	fmt.Println()
	if x%2 == 0 {
		for z = 0; z < y; z++ {
			fmt.Printf("%d ", x)
			x += 2
		}
	} else {
		fmt.Print("O PRIMEIRO NUMERO NAO E PAR\n")
	}
}
