// cf143

package main

import (
	"fmt"
)

func main(){
   var x interface{} = 7

   switch x.(type){
   case bool:
      fmt.Println("bool")
   case int:
      fmt.Println("int")
   case float32, float64:
      fmt.Println("float")
   }    

   fmt.Println("--------")
   
   // typeアサーションした変数を使うパターン
   switch v := x.(type){
   case bool:
      fmt.Println("bool", v)
   case int:
      fmt.Println("int", v)
   case float32, float64:
      fmt.Println("float", v)
   }
}
