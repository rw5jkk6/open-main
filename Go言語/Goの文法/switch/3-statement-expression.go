package main

import (
	"fmt"
)

func main(){
	// switch 簡易文; 式
	switch s := "autumn"; s{
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
