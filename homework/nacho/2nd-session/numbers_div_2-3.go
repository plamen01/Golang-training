// https://play.golang.org/p/rvyRa6seu9

package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		if i % 2 == 0 && i % 3 == 0 {
			fmt.Println(i, "is divisible by both")
		} else if i % 2 == 0 {
			fmt.Println(i, "is divisible by 2")
		} else if i % 3 == 0 {
			fmt.Println(i, "is divisible by 3")
		}
	}
}
