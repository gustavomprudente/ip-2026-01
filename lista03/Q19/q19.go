package main

import f "fmt"

func main () {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if j > i {
				f.Printf("[%d, %d] ", i, j)
			}
		}
	}
}