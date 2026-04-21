package main

import "fmt"

func main() {
	var n int

	fmt.Printf("Digite um número para transformá-lo em binário: ")
	fmt.Scan(&n)

	seq := []int{}

	for {
		if n%2 == 0 {
			seq = append(seq, 0)
			n /= 2
		} else {
			seq = append(seq, 1)
			n /= 2
		}
		if n < 1 {
			break
		}
	}
	i := 0
	j := len(seq) - 1

	for i < j {
		seq[i], seq[j] = seq[j], seq[i]
		i++
		j--
	}

	fmt.Printf("O número em binário é: %v\n", seq)
}