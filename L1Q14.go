package main

import (
	"fmt"
	"math"
)

func main(){
	var h, ar, ab, v float64
	fmt.Printf("DIGITE, RESPECTIVAMENTE, A ALTURA DA PIRÂMIDE E A ARESTA DO HEXÁGONO QUE FORMA A BASE DA PIRÂMIDE:\n")
	fmt.Scan(&h, &ar)
	ab = (3 * math.Pow(ar, 2) * math.Pow(3, 0.5)) / 2
	v = ab * h / 3
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", v )
}