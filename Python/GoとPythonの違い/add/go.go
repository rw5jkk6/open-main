package main

import (
	"fmt"
)

func add(n int, m int) int{
   return n + m
}

func calc_multi_return(n int, m int) (int ,int){
   a := n / m
   b := n % m
   return a ,b
}

func main(){
   num := add(1, 2)
   fmt.Println(num)

   num1, num2 := calc_multi_return(3, 2)
   fmt.Printf("割り算は%d 剰余は%d\n",num1, num2)
}
