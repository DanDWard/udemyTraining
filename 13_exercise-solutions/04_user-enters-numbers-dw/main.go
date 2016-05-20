package main

import "fmt"

func main()  {

	var lNumber int
	var sNumber int

	fmt.Print("Enter a large number:")
	fmt.Scan(&lNumber)
	fmt.Print("Enter a small number:")
	fmt.Scan(&sNumber)
	fmt.Println("Remainder of",lNumber," divided by",sNumber," is", lNumber%sNumber)

	//var numOne int
	//var numTwo int
	//fmt.Print("Please enter a large number: ")
	//fmt.Scan(&numOne)
	//fmt.Print("Please enter a smaller number: ")
	//fmt.Scan(&numTwo)
	//fmt.Println(numOne, "%", numTwo, " = ", numOne%numTwo)




}