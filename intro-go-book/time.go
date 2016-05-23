package main

import (
	"time"
	"fmt"
)


func main(){

	defer fmt.Println("Now is executed the defer func")

	for i:=0; i <= 10; i++ {

		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Current time is", time.Now())

	}
}
