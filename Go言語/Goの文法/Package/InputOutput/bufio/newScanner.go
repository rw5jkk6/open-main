package main

import (
   "os"
	"bufio"
	"fmt"
)

func main(){
   fp, err := os.Open("test.go")
   if err != nil{
      fmt.Println(err)
   }
   defer fp.Close()
   scanner := bufio.NewScanner(fp)
   for scanner.Scan(){
      fmt.Println(scanner.Text())
   }
}
