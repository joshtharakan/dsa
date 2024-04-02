package main

import "fmt"

func calculateSIPReturns(principal float64, initialNAV float64, rate float64, years int) float64 {

	nav := initialNAV
	total := 0.0
	for i := 1; i <= years; i++ {
		units := principal / nav
		total += units
		nav = nav * (1 + rate)
	}

	fmt.Println("Total units purchased: ", total)
	fmt.Println("Final NAV: ", nav)
	return total * nav
}

func main() {
	fmt.Println(calculateSIPReturns(50000, 10, 0.15, 5))

}
