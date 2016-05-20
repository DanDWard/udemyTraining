package main

import "fmt"

func main() {

	sum := 1

	for i := 1; i<=100; i++ {
		if i%3 == 0 {
			sum = sum + i
		} else if i%5 ==0 {
			sum = sum + i
		}

	}
	fmt.Println(sum)
}