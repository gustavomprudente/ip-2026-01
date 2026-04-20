package main 

import f "fmt"

func main () {
	var n int
	var t bool

	f.Printf("Digite um número para saber se ele é triangular.\n")
	f.Scan(&n)

	if n < 0 {
		f.Printf("Dígito inválido.\n")
		return
	}
	
	for i := 0; i * (i+1) * (i+2) <= n; i++ {
		t = false
		if i * (i+1) * (i+2) == n {
			t = true
		}
	}
	switch t {
		case true:
			f.Printf("%d é triangular.\n", n)
		default:
			f.Printf("%d não é triangular.\n", n)	
		}
}