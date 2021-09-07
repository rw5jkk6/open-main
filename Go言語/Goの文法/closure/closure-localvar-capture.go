package main

import (
	"fmt"
)

func main(){
	var closure func(int)
	{
		closure = func(a int){
			fmt.Println(a * 2)
		}
	}
	closure(2) // 4
}
