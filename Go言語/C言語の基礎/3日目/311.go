package main

import (
	"fmt"
)

func main(){
   // 下の２通りのどっちでもいい

   //d := []float64{1.2, 3.7, 4.1}
   var d [3]float64
   d[0]=1.2
   d[1]=3.7
   d[2]=4.1


   var sum , avg float64
   sum = 0
   for i:=0; i < 3; i++{
      fmt.Printf("%f ", d[i])
      sum += d[i]
   }
   fmt.Println("")
   avg = sum / 3.0
   fmt.Printf("合計値:%f\n", sum)
   fmt.Printf("平均値:%f\n", avg)
}

// 1.200000 3.700000 4.100000 
// 合計値:9.000000
// 平均値:3.000000
