package main

import (
	"fmt"
	"math/rand"
	"sort"
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

   sort.Ints(a)
   
   fmt.Println(a)
}
