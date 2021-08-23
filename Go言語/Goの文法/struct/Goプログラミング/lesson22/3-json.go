package main

import (
	"encoding/json"
	"fmt"
)


type pl struct{
   Name string
   Npl int
}

func diff2(n pl, s pl){
   d := s.Npl - n.Npl
   fmt.Printf("%sと%sの差は%dです。\n", n.Name, s.Name, d)
}


func main(){
   
   n := pl{Name: "nak", Npl: 3}
   s := pl{Name: "ser", Npl: 20}
   diff2(n, s)

   bytes ,_ := json.Marshal(n)
   fmt.Println(string(bytes))   // {"Name":"nak","Npl":3}
}
