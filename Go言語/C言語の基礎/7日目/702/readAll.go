package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main(){
   file, err :=os.Open("test.txt")
   if err != nil{
      log.Fatal(err)
   }

   content, err := ioutil.ReadAll(file)
   if err != nil{
      log.Fatal(err)
   }
   fmt.Println(string(content))
}

// 一度ファイルをオープンにしてから読み込む
