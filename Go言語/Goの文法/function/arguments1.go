package main

import (
	"fmt"
)

func callBack(i int, f func(int)int){
	fmt.Println(f(i))
}

func add10(j int) int{
	return j + 10
}

func main(){
	callBack(4, add10)
}
