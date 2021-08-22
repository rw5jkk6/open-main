package main

import (
	"fmt"
)

type human interface {
	say()
	walk()
}

type otokonomusume struct {
	name string
}

func (s otokonomusume) say() {
	fmt.Println("せりにゃん好き")
}

func (s otokonomusume) walk() {
	fmt.Println("せりにゃんに会いに行く")
}

func main() {
	n := otokonomusume{"hitoshi"}
   n.say()
   n.walk()
}
