package main
 
import (
   "fmt"
)

func main(){
   var num int
   fmt.Println("1~6の値を入力してください")

   fmt.Scanf("%d", &num)
   
   switch num{
   case 1:
      fallthrough
   case 3:
      fallthrough
   case 5:
      fmt.Println("奇数です")
   case 2:
      fallthrough
   case 4:
      fallthrough
   case 6:
      fmt.Println("偶数です")
   default:
      fmt.Println("範囲外です")
   }
}
