package main

import "fmt"

func main () {

	var r1, r2 []int
	var soma int
	n1 := make([]int, 10)
	n2 := make([]int, 5)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d° número do 1° vetor: ", i+1)
		fmt.Scan(&n1[i])
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Printf("%d° número do 2° vetor: ", i+1)
		fmt.Scan(&n2[i])
		soma += n2[i]
	}

	for i := 0; i < 10; i++ {
		if n1[i]%2 == 0 {
			r1 = append(r1, n1[i]+soma)
		} else {
			r2 = append(r2, n1[i]+soma)
		}
	}

	fmt.Printf("Vetor resultante 1: %d\n", r1)
	fmt.Printf("Vetor resultante 2: %d\n", r2)
}	