package main

import f "fmt"

func main() {
    var S, n float64
    n = 1
    for i := 15; i > 0; i-- {
        if i%2 != 0 {
            S += n / float64(i*i)
        } else {
            S -= n / float64(i*i)
        }
        n *= 2
    }
    f.Printf("S = %.2f\n", S)
}