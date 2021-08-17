package main

import (
	"fmt"
	"sort"
)

func main(){
   a :=[]string{"banana", "apple", "orange", "lemon", "pineapple"}
   fmt.Println(a)
   sort.StringSlice(a).Sort()
   fmt.Println(a)
}

// [banana apple orange lemon pineapple]
// [apple banana lemon orange pineapple]
