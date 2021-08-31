/*
無名関数を変数に入れている
こういうのを関数リテラルと呼ぶ
*/

package main
 
import (
   "fmt"
)

func main(){
	f := func(n int, m int)int{
		return n + m
	}
	r := f(1,2)
	fmt.Println(r)
}
