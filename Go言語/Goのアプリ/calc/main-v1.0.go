package main

import (
	"fmt"
	"math"
)

const (
   ten2 int = iota
   ten16 
   pow
   exit
)

func main(){
   var num int
   var ten2_num int
   var ten16_num int
   var pow_num float64

   for {

      fmt.Println("何を計算しますか？")
      fmt.Println("0 <-10進数を 2進数へ")
      fmt.Println("1 <-10進数を16進数へ")
      fmt.Println("2 <-2の累乗を計算")
      fmt.Println("3 <-終わる")
      fmt.Scanf("%d",&num)

      switch num{
      case ten2:
         fmt.Scanf("%d",&ten2_num)
         fmt.Println("---------------")
         fmt.Printf("%b\n", ten2_num)
         fmt.Println("---------------")

      case ten16:
         fmt.Scanln(&ten16_num)
         fmt.Println("---------------")
         fmt.Printf("%x\n", ten16_num)
         fmt.Println("---------------")

      case pow:
         fmt.Scanln(&pow_num)
         fmt.Println("---------------")
         fmt.Printf("%v\n", math.Pow(2, pow_num))
         fmt.Println("---------------")

      case exit:
         return
      }
   }
}
