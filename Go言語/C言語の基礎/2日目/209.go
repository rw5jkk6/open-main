package main

import (
	"fmt"
)

func main(){
	var a int
	fmt.Printf("数値を入力: ")
	fmt.Scanf("%d", &a)
	if a > 0{
		fmt.Println("入力した数は正数です")
	}else{
		fmt.Println("入力した数は正の数ではありません")
	}
}
