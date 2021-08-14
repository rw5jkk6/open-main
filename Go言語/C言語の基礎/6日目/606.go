// Goではポインタ演算はできないので、
// show関数は作れない。unsafe.Pointerを使うと作れる

package main

import (
	"fmt"
)

func main(){
   s :=[3]string{"Taro", "Hanako", "Tom"}
   fmt.Println("** Sの配列として表示 **")
   for i:=0; i < 3; i++{
      fmt.Println(s[i])
   }    

}
