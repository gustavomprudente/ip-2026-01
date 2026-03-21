package main

import "fmt"

func main(){
	var n1, n2, n3 float32
	fmt.Printf("Digite 3 números:") 

	fmt.Scan(&n1, &n2, &n3)

	media := (n1 + n2 + n3)/3

	fmt.Printf("MEDIA = %.2f\n", media)

	if media >= 6 {
		fmt.Println("APROVADO\n")
	} else {
	fmt.Println("REPROVADO\n")
	}
	
}
