package main

import (
	"fmt"
)

type A struct{
	i int
}

func main(){
	list := []A{{i:1}, {i:2},{i:3}}

	pList := make([]*A, 0, len(list))
	
	// for _, v := range list{
	// 	pList = append(pList, &v)
	// }

	// 3
	// 3
	// 3

	for i := range list{
		pList = append(pList, &list[i])
	}

	// 1
	// 2
	// 3

	for _, v := range pList{
		fmt.Println(v.i)
	}
}