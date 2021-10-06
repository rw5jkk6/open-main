package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
   os.Mkdir("testDir", 0755)
   os.Create("testDir/a.txt")
   

   f, err := os.Open(".")
   if err != nil{
      log.Fatal(err)
   }

   // openさせたら必ずcloseする
   defer f.Close()

   fis, err := f.Readdir(0)
   for _ , fi := range fis{
      if fi.IsDir(){
         fmt.Println(fi.Name())
      }
   }
}
