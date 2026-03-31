package main 

import "fmt"

func main () {
	var n [10]float64

	for i:= 0; i < 10; i++ {
		fmt.Printf("Digite o %d° número: ", i+1)
		fmt.Scan(&n[i])
	}
	fmt.Println("Os números em ordem inversa que foram inseridos ficam:")

	for i := 9; i >= 0; i--  {
	fmt.Printf("%.2f\n",n[i])
	}
}
