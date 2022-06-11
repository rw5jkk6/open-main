package main

import (
	"fmt"
	"fizzbuzz/lib"
)

func main(){
	for i := 1; i < 1000; i++{
		n := lib.FizzBuzz(i)
		fmt.Println(n)
	}
}