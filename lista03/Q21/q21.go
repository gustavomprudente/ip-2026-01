package main

import f "fmt"

func main () {
	var b, e, r int

	f.Printf("base: ")
	f.Scan(&b)
	f.Printf("expoente: ")
	f.Scan(&e)
	r = 1
	for i := 0; i < e; i++ {
		r *= b
	} 
	f.Printf("%d^%d = %d\n", b, e, r)
}