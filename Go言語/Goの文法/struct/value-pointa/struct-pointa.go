package main

import (
	"fmt"
)

type T struct{
	Number int
	Text string
}

type Container struct{
	V *T
}

func main(){
	c := Container{V: &T{}}
	v := c.V
	v.Number = 1
	fmt.Println(c.V.Number) // 1
	c.V.Text = "Hello"
	fmt.Println(c.V.Text)   // Hello
}