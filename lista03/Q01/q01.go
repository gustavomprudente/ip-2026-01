package main

import "fmt"

func main () {
	
	var b, e int
	
	fmt.Printf("Base: ")	
	fmt.Scan(&b)
	fmt.Printf("Expoente: ")
	fmt.Scan(&e)
	r := 1
	if e >= 0 {
		for i := 0; i < e; i++ {
			r *= b 
		} 
	} else {
		fmt.Printf("Dígito inválido.\n")
		return
	}
	
	fmt.Printf("Resultado: %d\n", r)
}