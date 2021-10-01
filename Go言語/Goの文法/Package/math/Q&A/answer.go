package main

import (
	"fmt"
	"math"
)

func main(){
   fmt.Println(math.Pi * 4.0)
   fmt.Println(math.Pow(2, 32))
   fmt.Println(math.Sqrt(2))
   fmt.Println(math.Max(3, 2))  // 3
   fmt.Println(math.Ceil(4.1))  // 5
　　　　　　　// 自然対数
   fmt.Println(math.Log(100))  // 4.605170185988092
   // 2を底とする
   fmt.Println(math.Log2(100))  // 6.643856189774724
}
