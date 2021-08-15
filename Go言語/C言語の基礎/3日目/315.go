package main

import (
	"fmt"
)

func main(){
   // var s1[4] = {'a', 'b'}  こういう書き方はできない
   var s2 string = "Hello world"
   var s3 string

   fmt.Println("文字列を入力してください")
   fmt.Scanf("%s",&s3)
   fmt.Printf("s2=%s\n", s2)
   fmt.Printf("s3=%s\n", s3)
}
