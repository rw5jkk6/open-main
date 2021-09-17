package main
 
import (
   "fmt"
)
func main(){
   a := 10
   b := 1.2345
   fmt.Printf("%4d  %.2f\n", a, b)    //   10  1.23
   fmt.Printf("%04d  %8.2f\n", a, b)  // 0010      1.23
   fmt.Printf("%v  %v\n", a, b)       // 10  1.2345
   fmt.Printf("%T  %T\n", a, b)       // int  float64
   fmt.Printf("%p\n", &a)             // 0xc00010c008
 
   var i int8 = 14
   fmt.Printf("%d\n", i)   // 14
   fmt.Printf("%b\n", i)   // 1110
   fmt.Printf("%x\n", i)   // e
}
