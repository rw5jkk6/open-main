package main

import (
	"fmt"
)

type human interface {
	say()
	walk()
}

// インターフェイスを満足させるには２つの
// メソッドを実装する必要がある
type otokonomusume struct {
	name string
}

func (s otokonomusume) say() {
	fmt.Println("せりにゃん好き")
}

func (s otokonomusume) walk() {
	fmt.Println("せりにゃんに会いに行く")
}

type onna struct{
   name string
}

func (o onna)say(){
   fmt.Println("女は好きではない")
}

func (o onna)walk(){
   fmt.Println("女の近くには行かない")
}

func main() {
   var n human
   n = onna{"hitoshi"}
   n.say()
   n.walk()
	n = otokonomusume{"hitoshi"}
   n.say()
   n.walk()
}
