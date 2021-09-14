package main

import (
	"fmt"
)

func main(){
	h := struct{
		name string
		pl int
	}{
		name: "Hitoshi",
		pl: 3,
	}

	fmt.Printf("%sのplは%dです\n", h.name, h.pl)
}