package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
   content, err := ioutil.ReadFile("test.go")
   if err != nil{
      fmt.Println(err)
   }
   fmt.Println(string(content))
}
