package main

import (
	"fmt"
)

func main(){
	name := []string{"satou", "suzuki", "yamada"}
	for i, v := range name{
		fmt.Println(i, v)
	}

	fmt.Println("----------------")

	for i := range name{
		fmt.Println(i)
	}

	fmt.Println("----------------")

	for _ , v := range name{
		fmt.Println(v)
	}
}