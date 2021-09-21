package main

import (
	"fmt"
)

func main(){
	var s interface{}
	s = "summer"
	switch s.(type){
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
