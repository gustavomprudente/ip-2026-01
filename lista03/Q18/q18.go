package main

import f "fmt"

func main () {
	for i := 0; i < 10; i++ {
		j := i
		f.Printf("[%d, %d] ", i, j)
	}
}