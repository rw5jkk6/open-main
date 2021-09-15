package main

import (
	"fmt"
)

func main(){
	var s interface{}
	// switch 簡易文; 式
	switch s = "summer"; s.(type){
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float64")
	case bool:
		fmt.Println("bool")
	default:
		return
	}
}
