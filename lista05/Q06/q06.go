package main

import "fmt"

func main () {
	num := make([]int, 100)
	
	for i := 100; i >= 1; i-- {
		num = append(num, i)
		if i == 1 {
			fmt.Printf("%d.\n", i)
			return
		}
		fmt.Printf("%d ", i)
	}
}