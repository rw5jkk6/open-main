package main

import (
	"fmt"
)

func main(){
   Loop:
      for {
         for{
            for{
               fmt.Println("start")
               break Loop
            }
            fmt.Println("one")
         }
         fmt.Println("two")
      }
      fmt.Println("finish")
}
// start
// finish
