package main

import (
	"fmt"
)

func main(){

   var a[9][9]int

   for i := 1; i < 10; i++{
      for j := 1; j < 10; j++{
         a[i-1][j-1] = i * j
      }
   }

   for i := 1; i < 10; i++{
      for j := 1; j < 10; j++{
         fmt.Printf("%d * %d = %2d  ", i, j,  a[i-1][j-1])
      }
      fmt.Println("")
   }
}
