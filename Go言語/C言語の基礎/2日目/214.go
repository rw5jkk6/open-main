package main
 
import (
   "fmt"
 
)


func main(){
   num := 0
   fmt.Printf("1~3の値を入力してください")
   fmt.Scanf("%d", &num)
   switch num{
   case 1:
      fmt.Println("one")
      break
   case 2:
      fmt.Println("two")
      break
   case 3:
      fmt.Println("three")
      break
   default:
      fmt.Println("不適切な値です")
      break
   }
}
