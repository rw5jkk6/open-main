package main

import (
	"fmt"
)

func main(){
   a := 10

   // p *int := &a だとエラーになる
   var p *int = &a  
   var pp **int = &p

   fmt.Println(p)
   fmt.Println(pp)
   fmt.Println(*pp)
   fmt.Println(**pp)

}

// 0xc000016048
// 0xc00000e028
// 0xc000016048
// 10
