package main

import "fmt"

func main () {
	var x, y int
	var tf, tc float64
	fmt.Println("Quantas temperaturas você quer converter?")
	fmt.Scan(&y)
	for x = 1; x <= y; x++ {
		fmt.Scan(&tf)
		tc = 5 * (tf - 32) / 9
		fmt.Printf("%f FAHRENHEIT EQUIVALE A %f CELSIUS\n", tf, tc)

	}

}