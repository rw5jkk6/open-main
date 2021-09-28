package main

import (
	"flag"
	"fmt"
)

func main(){
   var name string
   flag.StringVar(&name, "n", "", "")
   flag.StringVar(&name, "name", "", "")
   flag.Parse()
   fmt.Printf("%sは年内にせりにゃんと結婚する\n", name)
}
