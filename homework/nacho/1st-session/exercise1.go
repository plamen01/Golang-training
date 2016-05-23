package main

import (
	"fmt"
	"strconv"
)

var float_num float64 = 0.56789345

func main() {
	s := strconv.FormatFloat(float_num, 'f', 4, 64)
	fmt.Println(s[0:5])
}
