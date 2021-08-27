package main
 
import (
   "fmt"
)

func addFactory(n int)func(int)int{
	return func(m int)int{
		return n + m
	}
}

func main(){
	add5 := addFactory(5)
	r := add5(10)
	fmt.Println(r)

}