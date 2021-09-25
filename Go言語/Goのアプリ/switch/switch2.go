package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
   serina int = iota
   akari 
   haruka 
)

func main(){
   var name string
   flag.StringVar(&name, "n", "", "")
   flag.Parse()
   
   rand.Seed(time.Now().UnixNano())
   switch rand.Intn(4) {
   case serina:
      fmt.Printf("%sは、芹那とあなるせっくすする\n", name)
   case akari: 
      fmt.Printf("%sは、あかりの棒を受け入れる\n", name)
   case haruka:
      fmt.Printf("%sは、はるかを遠くから見る\n", name)
   default:
      fmt.Printf("%sは、犬で妥協する\n", name)
   }
}
