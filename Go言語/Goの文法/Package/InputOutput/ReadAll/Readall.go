package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
   f, err := os.Open("test.go")
   
   defer f.Close()
   
   if err != nil{
      fmt.Println(err)
   }

   content, err := ioutil.ReadAll(f)
   if err != nil{
      fmt.Println(err)
   }

   fmt.Println(string(content))

}
