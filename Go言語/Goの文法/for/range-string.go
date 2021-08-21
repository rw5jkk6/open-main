package main

import (
	"fmt"
)

func main(){
   for i, r := range "ABC"{
      fmt.Printf("[%d] -> %d\n", i, r)
   }   
}

// [0] -> 65
// [1] -> 66
// [2] -> 67