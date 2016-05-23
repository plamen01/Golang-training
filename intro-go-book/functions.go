package main

import (
	"fmt"
)

func add(args ...int) int {

	tot := 0

	for _, i := range args {
		tot += i
	
	}
	
	return tot


}

func main() {
	fmt.Println(add(1,2,3,4,5,6,7,8,9))
	
}
