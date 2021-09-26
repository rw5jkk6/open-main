package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct{
   Name     string
   Pl       int
   BeastPl int
}


func main(){
   src := `
{  
   "Name":"Hitoshi",
   "Pl":3,
   "BeastPl":5
}
`
   u := new(User)

   // json から　struct に変換
   err := json.Unmarshal([]byte(src), u)
   if err != nil{
      log.Fatal(err)
   }
   fmt.Printf("%+v\n", u)
}
