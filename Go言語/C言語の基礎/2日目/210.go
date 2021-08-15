package main
 
import (
   "fmt"
)

func main(){
   var num int
  
   fmt.Println("1~3の値を入力してください")

   fmt.Scanf("%d", &num)
   if num == 1{
      fmt.Println("one")
   }else if num == 2{
      fmt.Println("two")
   }else if num == 3{
      fmt.Println("three")
   }else{
      fmt.Println("不適切な値です")
   }
}
