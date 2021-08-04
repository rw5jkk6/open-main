package main

import (
	"fmt"
)

func swap(num1 *int, num2 *int){
	tmp := *num1
	*num1 = *num2
	*num2 = tmp
}

func main(){
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println("a:",a,"b:",b)	
}

// a: 2 b: 1
