// https://play.golang.org/p/6HYjCER8l9

package main

import (
	"fmt"
)

func main() {
	for a :=1 ; a <= 10; a++ {
		for b :=1 ; b <= 10; b++ {
			fmt.Println(a ,"*",b,"=", a*b)
		}
		fmt.Println("-----------")
	}
}
