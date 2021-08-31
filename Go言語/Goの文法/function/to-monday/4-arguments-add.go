/*
関数の引数に関数を入れることで、
コードに柔軟性を持たせている。
例えば、addをsub(引き算)とかに
できる
*/

package main

import (
	"fmt"
)

func add(n int, m int)int{
	return n + m
}

func calc(a int, b int, f func(int, int)int){
	c:=f(a, b)
	fmt.Println(c)
}

func main(){
	calc(1, 2, add)
}
