package main

import (
	"fmt"
)

type Vertex struct{
	X, Y int
}

func Area(v Vertex) int{
	return v.X * v.Y
}


func main(){
	v := Vertex{X: 3, Y: 4}
	n := Area(v)
	fmt.Println(n) // 12
}
