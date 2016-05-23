package main

import (
	"fmt"
)


func main(){
	x := [5]float64{ 2, 4, 5, 6, 6 }

	var total float64 = 0
	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))

}
