package main

import (
	"fmt"
	"os"
)

func main(){
   f, err := os.Open("test.go")
   if  err != nil{
      fmt.Println(err)
   }
   defer f.Close()

   bs := make([]byte, 1024)
   f.Read(bs)
   fmt.Println(string(bs))
}
