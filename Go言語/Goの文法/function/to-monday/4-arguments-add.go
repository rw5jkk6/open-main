package main

import (
	"fmt"
)

func add(n int, m int)int{
	return n + m
}

func calc(a int, b int, f func(int, int)int){
	c:=f(a, b)
	fmt.Println(c)
}

func main(){
	calc(1, 2, add)
}