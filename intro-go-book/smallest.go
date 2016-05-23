package main

import (
	"fmt"
)

func main(){
	list := []int {
		24,32,54,12,
		45,232,43,56,
		652,32,34,17,
		221,23,523,134,
	}


	min := list[0]


	for _, check := range list {
		if check < min {
			min = check
		}

	}

	fmt.Println(min)

}
