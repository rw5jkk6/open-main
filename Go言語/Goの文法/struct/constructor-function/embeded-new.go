/*
ポイント
＊コンストラクター関数
＊埋め込みによる継承みたいなことができる
＊フィールドを転送できる
*/

package main
 
import (
   "fmt"
)

type Vertex struct{
   x, y int
}

type Vertex3D struct{
   Vertex
   z int
}

func (v Vertex3D)Area3D() int{
   return v.x * v.y * v.z
}

func New(x, y, z int) *Vertex3D{
   return &Vertex3D{
      Vertex{x, y}, z}
}

func main(){
   v := New(2, 3, 4)
   fmt.Println(v.Area3D())
}

// 24
