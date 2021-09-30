package main

import (
	"fmt"
	"time"
)

func main(){
   // time.Secondがないと一瞬で出力する
   time.Sleep(5*time.Second)
   fmt.Println(time.Second)  // 1s
   fmt.Println("Hello")
}
