package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)


func main(){
   name := os.Args[1]
   
   rand.Seed(time.Now().UnixNano())
   r := rand.Intn(4)
   
   if r == 0 {
      fmt.Printf("%sは、芹那とあなるせっくすする\n", name)
   }else if r == 1{ 
      fmt.Printf("%sは、あかりの棒を受け入れる\n", name)
   }else if r == 2{
      fmt.Printf("%sは、はるかを遠くから見る\n", name)
   }else{
      fmt.Printf("%sは、犬で妥協する\n", name)
   }
}
