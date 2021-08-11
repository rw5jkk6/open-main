package main

import (
	"fmt"
)

// プロパティ変数でキャプチャする
func addProp()func(){
    n := 0
    return func(){
        n = n+1
        fmt.Println(n)
    }

}


// 引数でキャプチャする 
func add(i int) func(int){
    return func(j int){
        fmt.Println(i + j)
    }
}

func main(){
    a := addProp()
    a()
    a()

    fmt.Println("-----")

    add10 := add(10)
    add10(5)


}