package main

import"fmt"

func main () {
	var a, b, c, d, det float32
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	det = (a * d) - (b * c)
	fmt.Printf("O VALOR DO DETERMINANTE E = %2.f\n", det)
}
