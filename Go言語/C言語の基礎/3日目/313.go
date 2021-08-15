package main
 
import (
   "fmt"
)

func main(){
   var a[3][4]int
   var m, n int

   for m=0; m < 3; m++{
      for n=0; n < 4; n++{
         a[m][n] = m * n
      }
   }
   for m=0; m < 3; m++{
      for n=0; n < 4; n++{
         fmt.Printf("%d*%d=%d ", m, n, a[m][n])
      }
      fmt.Println("")
   }
}

// 0*0=0 0*1=0 0*2=0 0*3=0 
// 1*0=0 1*1=1 1*2=2 1*3=3 
// 2*0=0 2*1=2 2*2=4 2*3=6 
