package main

import (
	"fmt"
	"strings"
)

type talker interface{
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

// 引数tはインターフェイスなので、laserもmartianも入る
func shout(t talker){
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

// 埋め込みでインターフェイスを満足させる
type starship struct{
	laser
}

func main(){
	shout(martian{}) // NACK NACK
	shout(laser(2))  // PEW PEW

	// laserを埋め込んだstructを使っている
	shout(starship{3})  // PEW PEW PEW
}