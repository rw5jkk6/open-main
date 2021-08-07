package main

import (
	"fmt"
)

type Vertex struct{
	X, Y int
}

func (v *Vertex)Scale(i int){
	v.X = v.X * i
	v.Y = v.Y * i
}


func main(){
	v := Vertex{X: 3, Y: 4}
	v.Scale(10)
	fmt.Println(v) // {30, 40}
}
