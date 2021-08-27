package main

import (
	"fmt"
)

func integers()func() int{
	var i int
	return func() int{
		i = i + 1
		return i
	}
}

func main(){
	i :=integers()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
}