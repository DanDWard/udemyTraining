package main

import "fmt"

func main() {

	lFib :=1
	hFib :=2
	tFib :=0

	sumEvenFib :=0

	for hFib <= 4000000 {
		if hFib%2 == 0 {
			sumEvenFib += hFib
		}
		tFib = lFib + hFib
		lFib = hFib
		hFib = tFib

	}

	fmt.Println("The total of the even Fibonacci numbers up to 4 million is:",sumEvenFib)

}
