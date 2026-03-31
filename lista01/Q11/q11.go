package main

import "fmt"

func main() {
	var x int
	fmt.Printf("Digite um número: ")
	fmt.Scan(&x)
	if (x % 3 == 0) && (x % 5 == 0) {
		fmt.Printf("O NUMERO E DIVISIVEL\n")
	}  else {
		fmt.Print("O NUMERO NAO E DIVISIVEL\n")
	}

}
