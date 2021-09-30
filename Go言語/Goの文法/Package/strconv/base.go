package main

import (
	"fmt"
	"strconv"
)

func main(){
   
   str := "100"
   var num int64= 200
   str2 := "zero"
   
   // int64 -> 文字列
   s := strconv.FormatInt(num, 10)
   fmt.Printf("%s\n", s)  // 200

   // 文字列 -> int
   n, err := strconv.ParseInt(str, 10, 0)
   if err != nil{
      fmt.Println(err)
   }
   fmt.Printf("%d\n", n) // 100

   // 文字列 -> ?
   n, err = strconv.ParseInt(str2, 10, 0)
   if err != nil{
      fmt.Println(err) // strconv.ParseInt: parsing "zero": invalid syntax
   }
   // falseなので0を返す
   fmt.Printf("%d\n", n)  // 0
}
