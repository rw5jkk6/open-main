package main

import (
	"fmt"
)

func otoko()string{
	return "Otoko"
}

func onna()string{
	return "Onna"
}

func main(){
	var otokonomusume func()string
	// otokonomusume := otoko これでもOK
	otokonomusume = otoko
	fmt.Println(otokonomusume())

	otokonomusume = onna
	fmt.Println(otokonomusume())
}
