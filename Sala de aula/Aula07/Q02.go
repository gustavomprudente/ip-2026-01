package main 

import f "fmt"

func main () {
	var n1, n2, n3 int
	
	f.Print("Digite o primeiro número: ")
	f.Scan(&n1)
	f.Print("Digite o segundo número: ")
	f.Scan(&n2)
	f.Print("Digite o terceiro número: ")
	f.Scan(&n3)

	maio := maior (n1, n2, n3)
	f.Printf("O maior número é o %d", maio)
}

func maior (n1, n2, n3 int) int {
	if (n1 > n2) && (n1 > n3) {
		return n1
	} else if (n2 > n3) && (n2 > n1) {
		return n2
	} else {
		return n3
	}
}
