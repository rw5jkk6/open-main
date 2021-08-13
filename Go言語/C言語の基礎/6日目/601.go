package main

import (
	"fmt"
)

const SIZE = 3

func main(){
   var p1 []int
   var p2 []float32

   p1 = make([]int, SIZE)
   p2 = make([]float32, SIZE)

   for i:=0; i < SIZE; i++{
      p1[i] = i
      p2[i] = float32(i) / 10.0
   }

   fmt.Println(p1) // [0 1 2]
   fmt.Println(p2) // [0 0.1 0.2]
}
