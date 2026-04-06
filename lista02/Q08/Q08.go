package main

import f "fmt"

func main () {
	
	var n int

	f.Println("Digite um número inteiro para verificar se ele está entre 20 e 90")
	f.Scan(&n)

	if n > 20 && n < 90 {
		f.Printf("%d está entre 20 e 90\n", n)
	} else {
		f.Printf("%d não está entre 20 e 90\n", n)
	}
}
