package main

import (
	"fmt"
)

func main(){
   /* 
   var s1[4] = {'a', 'b', 'c', '¥0'}  
   こういう書き方はできない
   これは下の書き方をするとcコンパイラが自動で文字列を
   charの配列にしてくれる特別な仕様のため   
   */
   var s2 string = "Hello world"
   var s3 string

   fmt.Println("文字列を入力してください")
   fmt.Scanf("%s",&s3)
   fmt.Printf("s2=%s\n", s2)
   fmt.Printf("s3=%s\n", s3)
}
