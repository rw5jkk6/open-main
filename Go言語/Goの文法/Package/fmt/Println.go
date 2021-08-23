package main

import (
	"fmt"
	"os"
)

func Println(a ...interface{})(n int, err error){
	return fmt.Fprintln(os.Stdout, a...)
}

func main(){
	h := "Hello world"
	Println(h)
}