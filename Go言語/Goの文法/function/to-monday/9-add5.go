package main
 
import (
   "fmt"
)

func add5(n int)func(int)int{
	r := n + 5
	return func(m int)int{
		return r + m
	}
}

func main(){
	add1 := add5(10)
	b := add1(1)
	fmt.Println(b)
}