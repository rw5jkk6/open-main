package main
 
import (
   "fmt"
)

func main(){

	// 32bitは \Uxxxx
	s := "\U0001F628"
	fmt.Println(s)

	fmt.Println("\U0001F600")
}