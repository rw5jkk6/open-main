package main

import (
   "time"
	"fmt"
	"math/rand"
	"os"
)

func main(){
   rand.Seed(time.Now().UnixNano())
   c := len(os.Args) - 1
   if c < 1{
      fmt.Fprintf(os.Stderr, "[usage] %s choice1 choice2...", os.Args[0])
      os.Exit(1)
   }
   fmt.Printf("%s\n",os.Args[rand.Intn(c) + 1])
}
