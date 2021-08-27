package main
 
import (
   "fmt"
)


func main(){
	f := func(n int, m int)int{
		return n + m
	}
	r := f(1,2)
	fmt.Println(r)
}