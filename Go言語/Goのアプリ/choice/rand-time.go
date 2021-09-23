package main

import (
	"fmt"
	"math/rand"
   "time"
)

func main(){
   // func (t Time) UnixNano() int64
   fmt.Println(time.Now().UnixNano())
   // 1632386598181448000

   // func Seed(seed int64)
   // seedを72で固定しているので同じ数字が出てくる
   rand.Seed(72)
   // rand.Intn(num) 0 ~ num の間の数字がランダムに出力する
   fmt.Println(rand.Intn(100)) // 20
}
