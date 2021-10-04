package main

import (
	"fmt"
	"time"
)

func main(){
   
   t := time.Now()
   fmt.Println(t)

   // Date()は８つの引数があるが全部埋める必要がある
   t1 := time.Date(2021,12,24,0,0,0,0,time.Local)
   fmt.Println(t1.Year(), t1.Month(), t1.Day())
   
   time.Sleep(3 * time.Second)
   fmt.Println("せりにゃん好き")
   
   t2 := t.AddDate(1, 2, 0)
   fmt.Println(t2)   
}
