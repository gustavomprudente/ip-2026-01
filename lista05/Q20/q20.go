package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {
	dado := make([]int, 20)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		val := rand.Intn(6)+1
		dado[i] = val
	}

	freq := make(map[int]int)
	for _, v := range dado {
		freq[v]++
	}

	fmt.Printf("\n           <Resultado>      <Frequência>\n")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d° Dado:        %d                 %d\n", i+1, dado[i], freq[dado[i]])
	}
}