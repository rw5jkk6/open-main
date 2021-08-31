/*
5の問題では引数でデータを保存していたが、ここでは関数内の
ローカル変数としてデータを保存している
*/


package main

import (
	"fmt"
)

func later() func(string)string{
	var store string
	return func(next string) string{
		s := store
		store = next
		return s
	}
}

func main(){
	f := later()
	fmt.Println(f("Golang"))
	fmt.Println(f("is"))
	fmt.Println(f("awesome"))
	

}
