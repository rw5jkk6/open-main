package main

import (
	"fmt"
)

func main(){
	s := "summer"
	switch s{
	case "spring":
		fmt.Println("spring")
	case "summer":
		fmt.Println("summer")
	case "autumn":
		fmt.Println("autumn")
	case "winter":
		fmt.Println("winter")
	default:
		return
	}
}
