package main

import f "fmt"

func main () {

	var vc, vv float64
	f.Println("Digite o valor de compra do produto:")
	f.Scan(&vc)

	switch {
		case (vc < 10) && (vc >= 0):
			vv = 1.7 * vc

		case (vc >= 10) && (vc < 30):
			vv =  1.5 * vc

		case (vc >= 30) && (vc < 50):	
			vv =  1.4 * vc 

		case (vc >= 50):
			vv = 1.3 * vc	
		default:
			f.Println("valor inválido\n")
	}
	f.Printf("O valor de venda é: %.2f\n", vv)
}
