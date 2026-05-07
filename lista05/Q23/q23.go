package main

import "fmt"

func restante(A []int) int {
	soma := 0
	for i := 0; i < len(A); i++ {
		soma += A[i]
	}
	return 24 - soma
}

func main() {
	var poltronaTipo, q, cont int

	pj := make([]int, 24)
	pc := make([]int, 24)

	for {

		if (restante(pj) + restante(pc)) == 0 {
			fmt.Printf("Não há mais poltronas disponíveis.\n")
			return
		}

		fmt.Printf("Você deseja poltrona no corredor ou na janela?\n1 - Corredor\n2 - Janela\n3 - Verificar se ainda têm poltronas\n4 - Encerrar programa.\n")
		fmt.Scan(&poltronaTipo)

		switch {
		case poltronaTipo == 1:
			if restante(pc) == 0 {
				fmt.Printf("Não há poltronas disponíveis no corredor. Tente achar alguma na janela.\n") // adicionado \n
				continue
			}

			fmt.Printf("Número das poltronas livres no corredor: ")
			for i := 0; i < 24; i++ {
				if pc[i] == 0 {
					fmt.Printf("%d ", i+1)
				}
			}
			fmt.Println()
			fmt.Printf("%d poltronas ainda estão disponíveis para a venda.\nQuantas você deseja comprar?\n", restante(pc))
			fmt.Scan(&q)

			if q > restante(pc) {
				fmt.Printf("Quantidade indisponível.\n")
				continue
			}

			cont = 0
			for i := 0; i < 24; i++ {
				if cont == q {
					break
				}
				if pc[i] == 0 {
					pc[i] = 1
					cont++
				}
			}
			fmt.Printf("Compra realizada com sucesso.\n")
			continue

		case poltronaTipo == 2:
			if restante(pj) == 0 {
				fmt.Printf("Não há poltronas disponíveis na janela. Tente achar alguma no corredor.\n")
				continue
			}

			fmt.Printf("Números das poltronas livres na janela: ")
			for i := 0; i < 24; i++ {
				if pj[i] == 0 {
					fmt.Printf("%d ", i+1)
				}
			}
			fmt.Printf("\n")

			fmt.Printf("%d poltronas ainda estão disponíveis para a venda.\nQuantas você deseja comprar?\n", restante(pj))
			fmt.Scan(&q)

			if q > restante(pj) {
				fmt.Printf("Quantidade indisponível.\n")
				continue
			}

			cont = 0
			for i := 0; i < 24; i++ {
				if cont == q {
					break
				}

				if pj[i] == 0 {
					pj[i] = 1
					cont++
				}
			}

			fmt.Printf("Compra realizada com sucesso.\n")
			continue

		case poltronaTipo == 3:
			fmt.Printf("Total de poltronas disponíveis: %d\nPoltronas disponíveis no corredor: %d\nPoltronas disponíveis na janela: %d\n", restante(pj)+restante(pc), restante(pc), restante(pj))
			continue

		case poltronaTipo == 4:
			if (restante(pj) + restante(pc)) == 0 {
				fmt.Printf("Não há mais poltronas disponíveis.\n")
			} else {
				fmt.Printf("Total de poltronas disponíveis: %d\nPoltronas disponíveis no corredor: %d\nPoltronas disponíveis na janela: %d\n", restante(pj)+restante(pc), restante(pc), restante(pj))
			}
			return

		default:
			fmt.Printf("Dígito inválido. Tente novamente.\n") 
			continue
		}
	}
}