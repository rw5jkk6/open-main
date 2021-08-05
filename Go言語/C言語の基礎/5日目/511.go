package main
 
import (
   "fmt"
)


func main(){
	
	a := 100
	b := 0.2
	s1 := fmt.Sprint(a)
	s2 := fmt.Sprint(b)
	fmt.Printf("%s, %T\n", s1, s1)
	fmt.Printf("%s, %T\n", s2, s2)
}

// 100, string
// 0.2, string
