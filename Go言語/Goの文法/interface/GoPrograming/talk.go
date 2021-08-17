package main

import (
	"fmt"
	"strings"
)

var t interface{
	talk() string
}

type martian struct{}

func (m martian)talk() string{
	return "nack nack"
}

type laser int

func (l laser)talk() string{
	return strings.Repeat("pew ", int(l))
}
func main(){
	/*
	変数tはtalk()string型を持っているので、martianもlaserの両方の
	型に入れることができる
	*/

	t = martian{}
	fmt.Println(t.talk())  // nack nack

	t = laser(3)
	fmt.Println(t.talk())  // pew pew pew 
}
