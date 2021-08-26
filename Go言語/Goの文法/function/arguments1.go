package main

import (
	"fmt"
)

func add10(i int, f func(int)int){
	result := f(i)
	fmt.Println(result)
}

func callBack(j int) int{
	return j + 10
}

func main(){
	add10(4, callBack)
}

// 14
