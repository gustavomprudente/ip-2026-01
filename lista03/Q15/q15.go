package main

import f "fmt"

func main () {
	var N int

	f.Printf("Até qual termo o programa deve gerar?\n")
	f.Scan(&N)
	for i := 1; i <= N; i++ {
		if i == N {
			f.Printf("%v.\n", i*i)
		} else {
			f.Printf("%v ", i*i)
		}
	}
}