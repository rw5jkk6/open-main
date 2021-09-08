package main

import (
	"fmt"
)

func sum(a ...int)(sum int){
	for _, v := range a{
		sum += v
	}
	return sum
}


func main(){
	// 配列の作成
	arr := [...]int{1, 2, 3}
	// スライスの作成
	arr2:=arr[:]
	// スライスの拡張
	// func append(s []T, vs ...T) []T
	arr2 = append(arr2, arr2...)
	s := sum(arr2...)
	fmt.Println(s)	
}

// 12
