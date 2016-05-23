package main

import (
	"fmt"
)

var array = []int {1 , 2 , 3, 4,}

func greatest() {
	n := array[0]
	for _, check := range array {
		if check > n {
			n = check
		}
	}
	fmt.Println(n)
}


func main() {

	greatest()

}
