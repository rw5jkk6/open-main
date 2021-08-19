package main
 
import (
   "fmt"
)

//整数値を背の高さの型にする
type tall int16
type np int16

// 背が１センチ伸びるGromメソッド
func (t tall)Grow(){
   t += 1
   fmt.Println(t)
}

func (n np)Grow(){
   n += 1
   fmt.Println(n)
}

func main(){
   var t tall = 175
   t.Grow() // 176

   var n np = 3
   n.Grow() // 4
}
