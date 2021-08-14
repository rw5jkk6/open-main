package main

import (
	"fmt"
	"time"
)

func main(){
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for{
		select{
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("Boom")
			return 
		default:
			fmt.Println("   ")
			// 下の消すとスペースが延々と続く
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//   
//   
// tick
   
   
// tick
   
   
// tick
   
   
// tick
   
   
// tick
// Boom