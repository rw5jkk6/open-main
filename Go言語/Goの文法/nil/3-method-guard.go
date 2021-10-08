package main

import (
	"fmt"
)

type person struct{
   age int
}

// func (p *person)inc(){
//    p.age++
// }

// 2-panic.goでは != だったので注意
// errorを出さないようにしている
func (p *person)inc(){
   if p == nil{
      return 
   }
   p.age++
   
   // これでもOK
   // if p != nil{
   //    p.age++
   // }
}

func main(){
   var p *person
   p.inc()
   fmt.Println(p) // <nil>
}
