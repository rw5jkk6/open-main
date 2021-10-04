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
   user := new(User)
   user.Name = "Hitoshi"
   user.Pl = 3
   user.BeastPl = 5

   bs, err := json.Marshal(user)
   if err != nil{
      log.Fatal(err)
   }

   fmt.Println(string(bs))
}
