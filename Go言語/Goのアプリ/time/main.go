package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main(){
   var name string
   flag.StringVar(&name, "n", "" , "")
   flag.Parse()

   rand.Seed(time.Now().UnixNano())
   r := rand.Intn(10)
   t := time.Now()
   t1 := t.AddDate(r, 0, 0)
   fmt.Printf("%sはせりにゃんと%d年に付き合って\n", name, t1.Year())
   t2 := t1.AddDate(1, 2, 0)
   fmt.Printf("%d年%d月に結婚する\n", t2.Year(), t2.Month())
}
