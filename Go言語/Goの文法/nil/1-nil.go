package main

import (
	"fmt"
)

func main(){
	var i int
	fmt.Println(i)  // 0

	var ii *int
	fmt.Println(ii)  // <nil>

	var n [2]int      
	fmt.Println(n)   // [0,0]
	
	var nn *[2]int
	fmt.Println(nn)   // <nil>

	var m []int      
	fmt.Println(m)   // []  <- <nil>　と等価らしい []はポインタみたいなものか
	
	var mm *[]int
	fmt.Println(mm)   // <nil>
	
	c := make([]int, 4, 5)
	fmt.Println(c)    // [0 0 0 0]

	var a interface{}
	fmt.Println(a)   // <nil>

	var aa *interface{}
	fmt.Println(aa)  // <nil>
}
