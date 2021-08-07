package main

import (
	"fmt"
)

type Vertex struct{
	X, Y int
}

func Scale(v *Vertex, i int){
	v.X = v.X * i
	v.Y = v.Y * i
}


func main(){
	v := Vertex{X: 3, Y: 4}
	Scale(&v, 10)
	fmt.Println(v) // {30, 40}
}
