package main
 
import (
   "fmt"
)

func main(){
   
   // 配列
   a := [3]int{1, 2, 3}
   aa := a
   aa[0]=5
   fmt.Println(a)   // [1 2 3]
   fmt.Println(aa)  // [5 2 3]

   // スライス
   b := []int{1, 2, 3}
   bb := b
   bb[0]=5
   fmt.Println(b)   // [5 2 3]
   fmt.Println(bb)  // [5 2 3]

   // スライスに要素の追加
   bb=append(bb, 4, 5, 6, 7)
   fmt.Println(bb)  // [5 2 3 4 5 6 7]
}
