package main

import (
	"fmt"
)


func main(){
   a := [...]int{1,2,3}
   b := a
   b[0]=4
   fmt.Println(a)  // [1 2 3]
   c := &a
   c[0]=4
   fmt.Println(*c) // [4 2 3]
   fmt.Println(a)  // [4 2 3]
}
