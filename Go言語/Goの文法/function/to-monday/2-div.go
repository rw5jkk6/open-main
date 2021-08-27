package main

import (
	"fmt"
)

func div(n int, m int)(int, int){
	a := n / m
	b := n % m
	return a, b
}

func main(){
	a, b := div(15, 2)
	fmt.Printf("商は%dで剰余は%d\n", a, b)
}