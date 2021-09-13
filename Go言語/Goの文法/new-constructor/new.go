package main

import (
	"fmt"
)

type Person struct{
	name string
	age int
}

func main(){
	p1 := new(Person)
	fmt.Println(p1)  // &{ 0}

	p2 := &Person{}
	fmt.Println(p2)  // &{ 0}
}