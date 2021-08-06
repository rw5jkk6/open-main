package main
 
import (
   "fmt"
 
)

func max_n(m, n int) int{
   if m > n{
      return m
   }
   return n
}

func min_n(m, n int) int{
   if m < n{
      return m
   }
   return n
}

func main(){
   // Cの書き方だと            int (*n)(int, int) = max_n
   // Goでは右の書き方でもいい　　n := max_n
   var n func(int, int) int = max_n
   
   a := 1
   b := 2
   fmt.Printf("%dと%dのうち、最大のものは、%d\n", a, b, n(a, b))
   n = min_n
   fmt.Printf("%dと%dのうち、最小のものは、%d\n", a, b, n(a, b))

}

// 1と2のうち、最大のものは、2
// 1と2のうち、最小のものは、1
