package main

import "fmt"

func main() {
	var (
		p      int
		dep    float64
		codigo int
	)

	cod := make([]int, 10)
	saldo := make([]float64, 10)
	for p = 0; p < 10; p++ {

		fmt.Printf("--%d pessoa--\n", p+1)

		for {
			fmt.Printf("Digite o código: ")
			fmt.Scan(&codigo)

			existe := false
			
			for i := 0; i < p; i++ {
				if codigo == cod[i] {
					existe = true
					break
				}
			}

			if existe {
				fmt.Printf("Código já existe. Tente novamente.\n")
				continue
			}
			cod[p] = codigo
			fmt.Printf("Digite o saldo: ")
			fmt.Scan(&saldo[p])
			break 
		}
	}

	for {

		fmt.Print(
			"\n1. Efetuar depósito\n" +
				"2. Efetuar saque\n" +
				"3. Consultar o ativo bancário\n" +
				"4. Finalizar o programa.\nR = ",
		)

		resp := 0
		fmt.Scan(&resp)

		switch resp {
		case 1:

			fmt.Print("Código da conta: ")
			fmt.Scan(&codigo)

			pos := -1 

			for i := 0; i < 10; i++ {
				if codigo == cod[i] {
					pos = i
					break
				}
			}

			if pos != -1 {

				fmt.Printf("Saldo da conta: R$ %.2f\nQuanto você deseja depositar?\n", saldo[pos])
				fmt.Scan(&dep)
				saldo[pos] = saldo[pos] + dep

			} else {
				fmt.Printf("Conta não encontrada.\n")
			}

		case 2:
			fmt.Print("Código da conta: ")
			fmt.Scan(&codigo)
			pos := -1

			for i := 0; i < 10; i++ {
				if codigo == cod[i] {
					pos = i
					break
				}
			}

			if pos != -1 {
				fmt.Printf("Saldo da conta: R$ %.2f\nQuanto você deseja sacar?\n", saldo[pos])
				fmt.Scan(&dep)

				if saldo[pos] >= dep {
					saldo[pos] = saldo[pos] - dep
				} else {
					fmt.Println("Saldo insuficiente.")
				}

			} else {
				fmt.Printf("Conta não encontrada.\n")
			}

		case 3:
			ativoBancario := 0.0

			for i := 0; i < 10; i++ {
				ativoBancario += saldo[i]
			}

			fmt.Printf("Ativo Bancário = %.2f\n", ativoBancario)
			
		case 4:
			fmt.Printf("\nPrograma finalizado.\n")
			return

		default:

			fmt.Println("Opção inválida.")
		}
	}
}