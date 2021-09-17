package main

import (
	"fmt"
)

func main(){
	// 四則演算
	fmt.Println(5/2)
	fmt.Println(5*2)
	fmt.Println(8%3)

	fmt.Println("------------")

	// bit演算  cf88
	fmt.Println(11 & 1)
	fmt.Println(2 & 1)
	fmt.Println(3 & 2) //10進数で返ってくる

	fmt.Println("------------")
	
	// シフト演算  cf91
	fmt.Println(3<<1)
	fmt.Println(3<<2)
	fmt.Println(10>>1)
	fmt.Println(10>>2)
	
	fmt.Println("------------")

	// 比較演算子
	fmt.Println(1==1)
	fmt.Println(5>=7)
	fmt.Println(6!=2)

}

/*
c言語ではbool値がないので、何を出力するか？
*/

// #include <stdio.h>
 
// int main(void){
//     int a = 2;
//     int b = 2;
    
//     printf("%d\n", a==b);
//     printf("%d\n", a!=b);
// }
