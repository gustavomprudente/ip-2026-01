package main

import "fmt"

func main () {
	var h, m, s, s2 int
	fmt.Println("Digite, respectivamente, as horas, minutos e segundos:")
	fmt.Scan(&h)
	fmt.Scan(&m)
	fmt.Scan(&s)
	s2 = (h * 3600) + (m * 60) + s
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", s2)
}