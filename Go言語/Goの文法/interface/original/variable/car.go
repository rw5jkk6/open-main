package main

import (
	"fmt"
)

type car interface{
   forWheelRun() 
}

type keicar struct{
   place string
}

func (k keicar)forWheelRun(){
   fmt.Printf("近くの%sに行く\n", k.place)
}

type sportscar struct{
   speed int
}

func (s sportscar)forWheelRun(){
   fmt.Printf("%dキロで走る\n", s.speed)
}

func main(){
   var c car = keicar{"コンビニ"}
   c.forWheelRun()
   c = sportscar{150}
   c.forWheelRun()
}
