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
		V T
	}
	
	var c Container
	v := c.V
	v.Number = 1
	fmt.Println(c.V.Number) // 0
	c.V.Text = "Hello"
	fmt.Println(c.V.Text)   // Hello
}