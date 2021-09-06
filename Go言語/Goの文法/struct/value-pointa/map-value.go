package main

import (
	"fmt"
)

type T struct{
	Number int
	Text string
}

func main(){
	c :=map[int]T{0: T{}}
	// cannot assign to struct field c[0].Number in map
	// c[0].Number = 1
	c[0] = T{Number: 1}
	fmt.Println(c[0].Number)  // 1
}