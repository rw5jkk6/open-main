package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(10)
	b := rand.Intn(10)
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}
