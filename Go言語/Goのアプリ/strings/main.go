package main

import (
	"fmt"
	"strings"
)

func main(){
   var name string
   fmt.Println("好きな男の娘は誰？")
   fmt.Scanf("%s",&name)
   str := strings.Repeat(name, 1000)
   fmt.Println(str)
}
