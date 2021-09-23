/*
go run test.go を使わずに
go build test.go を使う
*/


package main

import (
	"fmt"
	"os"
)

func main(){
   c := os.Args
   // ./test 1 2 3 4
   fmt.Println(c)  // [./test 1 2 3 4]

   // len()関数はスライスの要素数を返す
   fmt.Println(len(os.Args)) // 5
   // スライスを添字で0番目を抜き出す
   fmt.Println(os.Args[0]) // ./test
}
