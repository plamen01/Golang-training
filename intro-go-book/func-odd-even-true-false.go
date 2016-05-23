package main

import (
	"fmt"
)

func half(n int) int {
	result := n / 2
	if result % 2 == 0 {
		fmt.Println(result, true)
	} else {
		fmt.Println(result, false) }

	return result
}
func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanln(&input)
	half(input)
}
