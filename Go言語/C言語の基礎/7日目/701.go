package main

import (
	"os"
)

func main(){
   f, _ := os.Create("test2.txt")
   defer f.Close()
   
   content := "Test\n"
   _, _ = f.Write([]byte(content))
}
