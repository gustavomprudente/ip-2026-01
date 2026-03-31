package main 

import "fmt"

func main () {

	var (
		nota [3]float64
		media, soma float64
	)
	for i := 0; i < 3; i++ {
		fmt.Printf ("Digite a nota %d: ", i+1)
		fmt.Scan(&nota[i])
		soma += nota[i]
	} 
	media = soma / 3
	fmt.Printf("Média: %.2f\n", media)

	for i := range nota {
		if nota[i] >= 6 {
			fmt.Printf("A %d° nota está acima da média (%.2f)\n", i + 1, nota[i])
		} else {
			fmt.Printf("A %d° nota não está acima da média (%.2f)\n", i+1, nota[i])
		}
	}
}
