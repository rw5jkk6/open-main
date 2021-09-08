package main

import (
	"fmt"
)

func sum(a ...int)(sum int){
	for _, v := range a{
		sum += v
	}
	return sum
}


func main(){
	arr := [...]int{1, 2, 3}
	arr2:=arr[:]
	s := sum(arr2...)
	fmt.Println(s)	
}
