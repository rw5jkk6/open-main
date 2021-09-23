/*
Newメソッドの戻り(error)の時にErrorメソッドが呼ばれる
func New(text string) error{
    return &errorString{text}
}
type errorSting struct{
    s string
}
func (e *errorString) Error() string{
    return e.s
}
*/


package main

import (
	"errors"
	"fmt"
	"os"
)

// Newは構造体

func divide(x, y int)(int, error){
	if y == 0{
		return 0, errors.New("divide by zero")
	}
	return x / y, nil
}

func main(){
	result, err := divide(3, 0)
	if err != nil{
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("result: %d\n", result)
}

// error: divide by zero
// exit status 1
