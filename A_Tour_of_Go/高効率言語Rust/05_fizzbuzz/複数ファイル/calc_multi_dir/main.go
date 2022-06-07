package main
 
import (
   "fmt"
   "culuc/lib"
   "culuc/lib/etc"
)
func main(){
	fmt.Println(lib.Add(1, 2))
	fmt.Println(lib.Sub(2, 1))
	fmt.Println(etc.Mul(2, 3))
	fmt.Println(etc.Div(4, 2))
}