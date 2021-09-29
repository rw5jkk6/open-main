package main

import (
	"fmt"
	"strings"
)

func main(){
   str := []string{"  Hitoshi", "serina", "akari  "}
   str2 := strings.Join(str, ",")
   fmt.Println(strings.Join(str, ","))
   fmt.Println(strings.Contains(str2, "serina"))
   fmt.Println(strings.Repeat("serina好き", 100))
   fmt.Println(strings.Replace(str2, "serina", "haruka", -1))
   fmt.Println(strings.TrimSpace(str2))
   fmt.Println(strings.ToUpper(str2))   
}
