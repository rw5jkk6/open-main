package main

import (
	"fmt"
)

func main(){
	
	s  := "ABC"
	// 文字列のコピー
	ss := s
	fmt.Println(ss)

	// 文字列の結合
	ss = s + "DEF"
	fmt.Println(ss)

	// 文字の数　これだけ関数
	le := len(ss)
	fmt.Printf("文字列の長さ:%d\n",le)
}

// ABC
// ABCDEF
// 文字列の長さ:6
