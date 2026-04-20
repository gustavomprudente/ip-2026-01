package main

import f "fmt"

func main () {
	var soma int

	for i := 1; i <= 20; i++ {
		f.Printf("%d ", i)
		soma += i
	}
	f.Printf("\nsoma: %d\n", soma)
}