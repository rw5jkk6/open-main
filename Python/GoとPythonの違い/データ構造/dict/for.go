package main

import (
	"fmt"
)


func main(){
   d := map[string]int{"apple": 2, "banana":3, "orange": 1} 

   for k1 := range d{
      fmt.Println(k1)
   }

   fmt.Println("*****************")

   for k2, v2 := range d{
      fmt.Println(k2, v2)
   }

   fmt.Println("*****************")

   for _, v3 := range d{
      fmt.Println(v3)
   }
}
