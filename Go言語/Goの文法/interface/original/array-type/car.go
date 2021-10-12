package main

import (
	"fmt"
)

type car interface{
   forWheelRun() 
}

type keicar struct{
   speed int
   place string
}

func (k keicar)forWheelRun(){
   fmt.Printf("近くの%sに行く\n", k.place)
}

type sportscar struct{
   speed int
   place string
}

func (s sportscar)forWheelRun(){
   fmt.Printf("%dキロで走る\n", s.speed)
}

func main(){
   k := keicar{place: "コンビニ"}
   s := sportscar{speed:150}
   array := []car{k, s}
   for _, v := range array{
      v.forWheelRun()
   }
}
