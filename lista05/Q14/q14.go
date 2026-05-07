package main

import "fmt"

func main () {

	var n1, n2 [10]int
	var r []int

	for i := 0; i < 10; i++ {
		fmt.Printf("%d° número do 1° vetor: ", i+1)
		fmt.Scan(&n1[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d° número do 2° vetor: ", i+1)
		fmt.Scan(&n2[i])
	}
	for i := 0; i < 10; i++ {
		r = append(r, n1[i])
	}
	for i := 0; i < 10; i++ {
		r = append(r, n2[i])
	}
	fmt.Printf("%d", r)
}