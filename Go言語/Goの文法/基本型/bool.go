package main
 
import (
   "fmt"
)
func main(){
   t, f := true, false
   fmt.Printf("%T %v %t\n", t, 1, t)   
   fmt.Printf("%T %v %t\n", f, f, f)
   
   fmt.Println(true && true)
   fmt.Println(true && false)
   
   fmt.Println(true || true)
   fmt.Println(true || false)
   
   fmt.Println(!true)
   fmt.Println(!false)
}

// bool 1 true
// bool false false
// true
// false
// true
// true
// false
// true
