package main

import (
	"fmt"
	"os"
)

func main(){
   // 2
   os.Create("test.txt")
   // 3
   os.Rename("test.txt", "test2.txt")
   // 4
   os.Remove("test2.txt")
   // 5
   dir, _ := os.Getwd()
   fmt.Println(dir)
   // 6
   os.Mkdir("testdir", 0775)
   // 7
   fmt.Println(os.Environ())
}
