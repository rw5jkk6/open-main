package main

import (
	"fmt"
)

type Test struct{
	name string
}

func main(){
	a := Test{name: "a"}
	fmt.Printf("%+v\n", a)

	var b Test
	b.name = "b"
	fmt.Printf("%+v\n", b)
	
	c := new(Test)
	c.name = "c"
	fmt.Printf("%+v\n", *c)
}

// {name:a}
// {name:b}
// {name:c}
