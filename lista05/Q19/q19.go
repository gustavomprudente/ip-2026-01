package main

import "fmt"

func main () {
	num := make([]int, 10)
	divis := make([]int, 5)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d° número do 1° vetor: ", i+1)
		fmt.Scan(&num[i])
	}
fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Printf("%d° número do 2° vetor: ", i+1)
		fmt.Scan(&divis[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("\nNúmero %d:\n", num[i])
		t := true

		for j := 0; j < 5; j++ {
			if num[i]%divis[j] == 0 {
				fmt.Printf("	Divisível por %d na posição %d\n",  divis[j], j)
				t = false
			}
		}

		if t {
			fmt.Printf("	Não é divisível\n")
		}
	}
}