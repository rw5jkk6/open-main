package main

import (
	"fmt"
)

func main(){
   
   var p *int // memoryは確保されていない
   *p = 1
   fmt.Println(*p)
}

// panic: runtime error: invalid memory address or nil pointer dereference
