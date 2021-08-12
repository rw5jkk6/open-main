//starting Go cf172

package main

import (
	"fmt"
)

func pow_array(a [3]int){
   for i, v := range a{
      a[i] = v * v
   }
   return
}

func pow_slice(a []int){
   for i, v := range a{
      a[i] = v * v
   }
   return
}

func main(){
   a := [3]int{1, 2, 3}
   pow_array(a)
   fmt.Println(a) // [1 2 3]

   b := []int{1, 2, 3}
   pow_slice(b)
   fmt.Println(b)  // [1 4 9]
}
