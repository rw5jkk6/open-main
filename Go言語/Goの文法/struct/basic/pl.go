package main

import (
	"fmt"
)

type Person struct{
	name string
	pl int
}

func main(){
	// h := Person{name: "Hitoshi", pl: 3}
	h := Person{}
	h.name = "Hitoshi"
	h.pl = 3
	fmt.Printf("%sのplは%dです\n", h.name , h.pl)
}