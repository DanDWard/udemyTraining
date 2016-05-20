package main

import "fmt"

func main() {
	h, even := half(10)
	fmt.Println(h, even)
}

func half(i int) (int, bool) {
	return i/2, i%2 ==0

}