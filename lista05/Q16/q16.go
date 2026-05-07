package main

import "fmt"

func main () {

	id := make([]int, 50)
	for i := 0; i < 50; i++ {
		fmt.Printf("%d° idade: ", i+1)
		fmt.Scan(&id[i])
	}

	q := make(map[int]int)
	for i := 0; i < 50; i++ {
		q[id[i]]++
	}

	moda := 0
	f := 0
	for i := 0; i < 50; i++ {
		if q[id[i]] > f {
			f = q[id[i]]
			moda = id[i] 
		}
	}
	fmt.Printf("Moda das idades: %d\n", moda, f)
}