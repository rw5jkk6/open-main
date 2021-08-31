/*
add5 := addFactory(5)でaddFactory()の仮引数に５が保存される
add5()の中には　n=5が保存されているので、足した15が出力できる
*/


package main
 
import (
   "fmt"
)

func addFactory(n int)func(int)int{
	return func(m int)int{
		return n + m
	}
}

func main(){
	add5 := addFactory(5)
	r := add5(10)
	fmt.Println(r)

}
