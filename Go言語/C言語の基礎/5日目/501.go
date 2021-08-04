package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var a int = 100
	var b float64 = 123.4
	var c float32 = 123.4
	var d rune = 'a'

	// 書式はc言語と同じ
	fmt.Printf("aの値は%d、大きさは%dbyte、アドレスは0x%x\n", a, unsafe.Sizeof(a), &a)
	fmt.Printf("bの値は%f、大きさは%dbyte、アドレスは0x%x\n", b, unsafe.Sizeof(b), &b)
	fmt.Printf("cの値は%f、大きさは%dbyte、アドレスは0x%x\n", c, unsafe.Sizeof(c), &c)
	fmt.Printf("dの値は%c、大きさは%dbyte、アドレスは0x%x\n", d, unsafe.Sizeof(d), &d)
}

// aの値は100、大きさは8byte、アドレスは0xc0000b2008
// bの値は123.400000、大きさは8byte、アドレスは0xc0000b2010
// cの値は123.400002、大きさは4byte、アドレスは0xc0000b2018
// dの値はa、大きさは4byte、アドレスは0xc0000b201c
