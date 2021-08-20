package main

import (
	"fmt"
)

func diff(n string, npl int, s string, spl int){
   d := spl - npl
   fmt.Printf("%sと%sの差は%dです。\n", n, s, d)
}


func main(){
   diff("nak", 3, "serina", 20)
}