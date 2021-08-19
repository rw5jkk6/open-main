package main
 
import (
   "fmt"
)

func sum(a ...int) int{
   n := 0
   for _, v := range a{
      n += v
   }
   return n
}

func main(){
   fmt.Println(sum(1))
   fmt.Println(sum(1, 2))
   fmt.Println(sum(1, 2, 3))
}
