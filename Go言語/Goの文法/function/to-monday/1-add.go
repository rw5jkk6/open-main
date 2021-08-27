package main
 
import (
   "fmt"
)

func add(n int, m int)int{
	return n + m
}

func main(){
	c := add(1, 2)
	fmt.Println(c)
}