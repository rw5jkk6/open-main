package main

import (
	"fmt"
)

func addNumFactory(i int) func(i int) int{
   return func(n int) int{
      return i + n
   }
}

func main(){
   add5 := addNumFactory(5)
   r := add5(10)
   fmt.Println(r)
   
}
