package main
 
import (
   "fmt"
)

func main(){
   var i *int
   fmt.Println(i) // <nil>

   //fmt.Println(*i) 
   // panicを起こす
   // panic: runtime error: invalid memory address or nil pointer dereference

   if i != nil{
      fmt.Println(*i)
   }
}
