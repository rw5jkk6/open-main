package main
 
import (
   "fmt"
)

type T struct{
	Number int
	Text string
}

func main(){
	type Container struct{
		V *T
	}
	
	// valueと異なり複合リテラルになっている
	c := Container{V: &T{}}
	v := c.V
	v.Number = 1
	fmt.Println(c.V.Number) // 1
	c.V.Text = "Hello"
	fmt.Println(c.V.Text)   // Hello
}