package main

import (
	"fmt"
)

// 戻り値を複数返せる
func calc(x, y int) (int, int) {
   return x + y, x - y
}

// 戻り値に変数を宣言できる
func calc2(price, item int) (result int){
   result = price * item
   return result
}


func main(){
   r1, r2 := calc(2, 1)
   fmt.Printf("r1=%d  r2=%d\n", r1, r2) // r1=3 r2=1

   r3 := calc2(100, 2)
   fmt.Println(r3)  // 200

   // 無名関数
   f := func(x int){
      fmt.Println(x)  // 3
   } 
   f(3)
}