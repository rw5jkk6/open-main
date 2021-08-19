package main

import (
	"fmt"
)

type Crier interface{
	Cry() string
}

type Duck struct{}

func (d Duck)Cry() string{
	return "Quack"
}

func main(){
	// 当然下のでも動く
	// var c Duck =Duck{}
	var c Crier = Duck{}
	fmt.Println(c.Cry())
}

// 変数cの型にインターフェイスを使う