package main

import "fmt"

func main() {
    var N, n1, n2 int
    var n []int

    fmt.Printf("1° número: ")
    fmt.Scan(&n1)
    fmt.Printf("2° número: ")
    fmt.Scan(&n2)
    n = append(n, n1, n2)

    fmt.Printf("Quantos números terá a sequência?\n")
    fmt.Scan(&N)

    for i := 2; i < N; i++ {
        var novo int
		if i%2 == 1 {
            novo = n[i-1] - n[i-2]
        } else {
            novo = n[i-1] + n[i-2]
        }
        n = append(n, novo)
    }

    fmt.Printf("%v\n", n)
}