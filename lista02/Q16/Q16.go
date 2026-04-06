package main 

import "fmt"

func main () {

	var d, m1, a int
	var m2 string

	fmt.Print("Digite o dia, mês e ano, respectvamente: ")
	fmt.Scan(&d, &m1, &a)
	switch m1 {
		case 1:
			m2 = "janeiro"
		case 2:
			m2 = "fevereiro"
		case 3:
			m2 = "março"
		case 4:
			m2 = "abril"	
		case 5:
			m2 = "maio"
		case 6:
			m2 = "junho"
		case 7:
			m2 = "julho"
		case 8:
			m2 = "agosto"
		case 9:
			m2 = "setembro"
		case 10:
			m2 = "outubro"
		case 11:
			m2 = "novembro"
		case 12:
			m2 = "dezembro"			 
	}

	fmt.Printf("%d de %s de %d\n", d, m2, a)
}
