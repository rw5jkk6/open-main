package main

import (
	"math/rand"
   "fmt"
   "time"
)

const SIZE = 4

func main(){
   var a = make([]int, SIZE)

   rand.Seed(time.Now().UnixNano())
   for i := 0; i < SIZE; i++{
      a[i] = rand.Intn(10)
   } 
   fmt.Println(a)

   for i:=0; i < SIZE; i++{
      for j := 0; j < SIZE - 1; j++{
         if a[j] > a[j+1]{
            swap(&a[j], &a[j+1])
         }
      }
   }
   fmt.Println(a)
}

func swap(a *int, b *int){
   tmp := *a
   *a = *b
   *b = tmp
}
