package main

import (
	"fmt"
)

type T struct{
	Number int
	Text string
}

func main(){
	t:=T{}
	fmt.Println(t)  // {0 }
	t2:=new(T)
	fmt.Println(t2) // &{0 }

}