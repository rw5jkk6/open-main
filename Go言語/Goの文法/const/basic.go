package main

import (
	"fmt"
)

const beast_mode = 1.5

const (
	serina = 20
	hitoshi = 3
)

func main(){
	fmt.Printf("芹那のビーストモードは%fです\n", serina*beast_mode)
	fmt.Printf("ヒトシのビーストモードは%fです\n", hitoshi*beast_mode)
}