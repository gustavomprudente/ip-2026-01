package main

import f "fmt"

func quad(A int)bool {
	for i := 0; i * i <= A; i++{
		if i * i == A {
			return true
		} 
	}
	return false
}

func main () {
	
	var (
		n int
		valor bool
		números []int
	)

	for {
		f.Printf("\n--Digite um número para verificar se ele é ou não quadrado perfeito--\n  (digite um número menor ou igual a zero para terminar o programa)\n")
		f.Scan(&n)
		if n <= 0 {
			break
		}

		números = append(números, n)
		valor = false
		valor = quad(n)
				
		switch valor {
		case true:
			f.Printf("%d é quadrado perfeito.\n", n)
		default:
			f.Printf("%d não é quadrado perfeito.\n", n)	
		}
	}

	f.Printf("Números informados: %v\n", números)
}