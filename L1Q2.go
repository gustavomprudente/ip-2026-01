package main

import "fmt"

func main() {
	var x, y, p int
	var r, pp, pg, pa, pc float64

	fmt.Print("Digite quantos jogos você quer analisar: ")
	fmt.Scan(&y)

	for x = 1 ; x <= y; x++ {
		fmt.Scan(&p, &pp, &pg, &pa, &pc)
		r = (float64(p) * (pp/100)) + 5 * (float64(p) * (pg/100)) + 10 * (float64(p) * (pa / 100)) + 20 * (float64(p) * (pc / 100))
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", x, r)
	}
}