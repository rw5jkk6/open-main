package main

import (
	"fmt"
	"time"
)

func sub(){
   time.Sleep(3 * time.Second)
   fmt.Println("Sub")
}

func main(){
   go sub()

   // 下のコードの５を２にしてみる
   time.Sleep(5 * time.Second)
   fmt.Println("Main")
}
