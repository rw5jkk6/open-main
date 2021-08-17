package main

import (
	"fmt"
)

func main(){
	
//    var i1 int
//    var d1 float32
//    d1 = 1.23
//    i1 = d1   //Goでは暗黙的な型変換ができない

   var i2, j2 int
   var d2, e2 float32

   j2 = 3

   d2 = 1.23

   i2 = int(d2)
   e2 = float32(j2)

   fmt.Printf("%d\n", i2)  // 1 
   fmt.Printf("%f\n", e2)  // 3.000000
}
