/*
5-addFactory.goの５を足しただけで、
基本的に同じ
*/

package main
 
import (
   "fmt"
)

func add5(n int)func(int)int{
	return func(m int)int{
		return n + 5 + m
	}
}

func main(){
	add1 := add5(10)
	b := add1(1)
	fmt.Println(b)
}
