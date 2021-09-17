package main
 
import (
   "fmt"
)

type pl int

func main(){
	var nak pl
	var seri int
	nak = 3
	seri = 20
	// 下のは型が異なるのでエラーになる
	// fmt.Println(nak + seri) 
	newnak := pl(seri) + nak
	fmt.Println(newnak) // 23
}
