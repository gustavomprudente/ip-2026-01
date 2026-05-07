package main

import "fmt"

func main () {
	num := make([]int, 100)
	imp := 1
	for i := 0; i < len(num); i++ { 
		num = append(num, imp)
		if i == 99 {
			fmt.Printf("%d.\n", imp)
			return
		}
		fmt.Printf("%d ", imp)
		imp += 2
	}
}