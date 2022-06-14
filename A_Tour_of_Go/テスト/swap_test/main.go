package main
 
import (
   "fmt"
   "swap_test/lib"
)

func main(){
	s := lib.S{X:1, Y:2}
	s.Swap()
	fmt.Println(s.X)
}