package main
 
import (
   "fmt"
)
func main(){
   a := 10
   b := 1.2345
   fmt.Printf("%4d  %.2f\n", a, b)
   fmt.Printf("%04d  %8.2f\n", a, b)
   fmt.Printf("%v  %v\n", a, b)
   fmt.Printf("%T  %T\n", a, b)
   fmt.Printf("%p\n", &a)
}

// 10  1.23
// 0010      1.23
// 10  1.2345
// int  float64
// 0xc00010c008
