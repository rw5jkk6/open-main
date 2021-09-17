package main

import (
	"fmt"
	"unsafe"
)

func main(){
	s := "Hello world , Hello Go"
	fmt.Println(unsafe.Sizeof(s))
}
// 16
