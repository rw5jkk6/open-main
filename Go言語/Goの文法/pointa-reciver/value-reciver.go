package main

import (
	"fmt"
)

type Vertex struct{
	X, Y int
}

func (v Vertex)Area() int{
	return v.X * v.Y
}

func main(){
	v := Vertex{X: 3, Y: 4}
	n := v.Area()
	fmt.Println(n) // 12
}