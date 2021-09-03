/*
関数の引数に関数を入れるので、callbackと呼ばれるものがある。
これは非同期通信の中で順に処理される時によく使われる
*/

package main

import (
	"fmt"
)

func add10(i int, f func(int)int){
	result := f(i)
	fmt.Println(result)
}

func callBack(j int) int{
	return j + 10
}

func main(){
	add10(4, callBack)
}

// 14
