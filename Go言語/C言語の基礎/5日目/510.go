package main

import (
	"fmt"
	"strconv"
)

func main(){
	var s1 string = "1000"
	var s2 string = "12.345"

	var a int
	var b float64

	a, _ = strconv.Atoi(s1)
	b, _ = strconv.ParseFloat(s2, 64)
	fmt.Printf("%d,  %f\n", a, b)
}
// 1000,  12.345000
