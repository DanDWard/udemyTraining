package main

import (
	"sort"
	"fmt"
)

func main () {
	s := []string{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(s)

	sort.Strings(s)

	fmt.Println(s)

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

}