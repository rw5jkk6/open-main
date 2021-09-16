/*
コンストラクタ関数は構造体を初期化するためのもの
コンストラクタ関数は関数めいの頭をnewをつける
ここでは２パターンある
*/


package main
 
import (
   "fmt"
)

type pl struct{
   name string
   npl int
} 
// パターン１
func newpl(name string, npl int) *pl{
   p := new(pl)
   p.name = name
   p.npl = npl
   return p
}

// パターン２
func newpl2(name string, npl int) *pl{
   return &pl{
      name: name,
      npl: npl,
   }
}

func diff2(n *pl, s *pl){
   d := s.npl - n.npl
   fmt.Printf("%sと%sの差は%dです。\n", n.name, s.name, d)
}

func main(){
   nak := newpl("nak", 3)
   ser := newpl2("seri", 20)

   diff2(nak, ser)
}

// nakとseriの差は17です。
