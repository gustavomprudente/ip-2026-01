package main 

import f "fmt"

func main () {
	var v, soma, maior, menor, contPar, par, impar int
	var n []int
	f.Printf("Digite os números (finalize a entrada de dados com a digitação do número 30.000):\n")
	
	for {
		f.Scan(&v)
		n = append(n, v)
		soma += v
		
		switch {
		case len(n) == 1:
			maior = v
			menor = v
		default:
			if v > maior {
				maior = v
			}
			if v < menor {
				menor = v
			}
		}

		if v%2 == 0 {
			par += v
			contPar++
		} else {
			impar++	
		}

		if v == 30000 {
			f.Printf("\nFim.\n")
			break
		}
		
	}
	m := soma / len(n)
	mPar := par / contPar
	pImpar := float64(impar) / float64(len(n)) * 100

	f.Printf("\nsoma: %d\n", soma)
	f.Printf("quantidade de números: %d\n", len(n))
	f.Printf("média: %d\n", m)
	f.Printf("maior número: %d\n", maior)
	f.Printf("menor número: %d\n", menor)
	f.Printf("média dos números pares: %d\n", mPar)
	f.Printf("percentagem dos números ímpares entre todos os números digitados: %.2f %%\n", pImpar)
}