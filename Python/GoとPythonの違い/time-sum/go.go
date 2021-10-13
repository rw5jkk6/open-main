package main
 
import (
   "fmt"
)

func f(n int)int{
   var sum int
   for i := 0; i < n; i++{
      sum += i 
   }
   return sum
}

func main(){
   s := f(2<<25)
   fmt.Println(s)
}

// 2251799780130816

// real    0m0.285s
// user    0m0.021s
// sys     0m0.003s
