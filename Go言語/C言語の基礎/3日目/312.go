package main
 
import (
   "fmt"
)


func main(){
	
	var n [5]int

   for i := 0; i < 10; i++{
      fmt.Printf("i=%d ", i)
      n[i] = i
   }
   fmt.Println("")
}

// i=0 i=1 i=2 i=3 i=4 i=5 panic: runtime error: index out of range [5] with length 5
