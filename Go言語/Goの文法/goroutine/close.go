package main

import (
	"fmt"
	"time"
)

func receive(name string, ch <- chan int){
	for {
		i, ok := <- ch
		if ok == false{
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func main(){
	ch := make(chan int, 2)

	go receive("1st goroutine", ch)
	go receive("2st goroutine", ch)
	go receive("3st goroutine", ch)

	i := 0
	for i < 10{
		ch <- i
		i++
	}
	close(ch)
	time.Sleep(3 * time.Second)
}

// 3st goroutine 0
// 1st goroutine 1
// 1st goroutine 3
// 1st goroutine 5
// 1st goroutine 6
// 1st goroutine 7
// 1st goroutine 8
// 1st goroutine 9
// 3st goroutine 2
// 1st goroutine is done.
// 3st goroutine is done.
// 2st goroutine 4
// 2st goroutine is done.