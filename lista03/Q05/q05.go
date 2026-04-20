package main

import f "fmt"

type Pessoa struct {
	idade int
	altura float64
	peso float64
}

func main () {
	var (
		id, continuar, q50, q20, q40 int
		a, p, soma float64
		pessoas []Pessoa
	) 

	for {
		f.Printf("Digite, respectivamente, a idade, a altura e o peso da pessoa:\n")
		f.Scan(&id, &a, &p)
		pes := Pessoa{
			idade: id,
			altura: a,
			peso: p,
		}
		pessoas = append(pessoas, pes)
		
		
		f.Printf("Deseja continuar adicionando pessoas?\n1 - Sim\nOutro valor diferente de 1 - Não.\n")
		f.Scan(&continuar)
		if continuar != 1 {
			break
		}
	}

	for i := 0; i < len(pessoas) - 1; i++ {
		if pessoas[i].idade > 50 {
			q50++
		}
	}

	for i := 0; i < len(pessoas) - 1; i++ {
		if pessoas[i].idade >= 10 && pessoas[i].idade >= 20 {
			soma += pessoas[i].altura
			q20++
		}
	}
	for i := 0; i < len(pessoas) - 1; i++ {
		if pessoas[i].peso < 40 {
			q40++
		}
	}
	m := soma / float64(q20)
	pinf := q40 / len(pessoas) * 100

	f.Printf("\n%d pessoas possuem idade superior a 50 anos\n", q50)
	f.Printf("%.2f é a média das alturas das pessoas com idade entre 10 e 20 anos\n", m)
	f.Printf("%d das pessoas analisadas possuem peso inferior a 40 quilos\n", pinf)
}