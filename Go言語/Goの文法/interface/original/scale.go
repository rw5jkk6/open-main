package main

import (
	"fmt"
)

type otoko interface{
   tinpo()
}
////////////////////////////////////////////
type Scale struct{
   size string
}

func (s Scale)tinpo(){
   fmt.Printf("チンポの大きさは%sです\n", s.size)
}
////////////////////////////////////////////

type Length struct{
   len int
}

func (l Length)tinpo(){
   fmt.Printf("チンポの長さは%dセンチです\n", l.len)
}
////////////////////////////////////////////

func Say(name string, o otoko){
   fmt.Printf("%sの", name)
   o.tinpo()
}

func main(){
   h := Scale{"ポークビッツ"}
   Say("Hitoshi", h)

   h2 := Length{4}
   Say("Hitoshi", h2)   


   s := Scale{"大根"}
   Say("Serina", s)

   s2 := Length{20}
   Say("Serina", s2)
}
