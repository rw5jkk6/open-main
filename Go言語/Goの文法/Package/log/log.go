package main

import (
	"fmt"
	"log"
)

func main(){
	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "test", "test") // <- ここでエラーがおこる
	log.Fatalf("%T %v", "test", "test2") // <- これは実行されない

	fmt.Println("ok?")
}

