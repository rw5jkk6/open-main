package main

import (
	"fmt"
)

func main(){
   var i2, j2 int
   var d2, e2 float32

   j2 = 3

   d2 = 1.23

   i2 = int(d2)
   e2 = float32(j2)

   fmt.Printf("%d\n", i2)  // 1 
   fmt.Printf("%f\n", e2)  // 3.000000
}
