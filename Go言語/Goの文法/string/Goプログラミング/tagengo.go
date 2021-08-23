package main

import (
	"fmt"
)

func main(){

	// 9-1
	var blank string
	fmt.Println(blank) // ""

	fmt.Println(`Hello \n world`)  // Hello \n world
	fmt.Println("Hello \nworld")   // Hello
	                               // world

	fmt.Println("A")    // A
	fmt.Println('A')    // 65

	// 9-3
	message := "serina"
	// message[0] ="a"  　errorになる
	fmt.Println(message) // serina

	// 9-4
	// ポインタ演算みたいなことができる
	c := 'a'
	c = c + 3
	fmt.Printf("%c\n", c)
}