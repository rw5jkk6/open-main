package main

import (
	"fmt"
)

func main(){
   var seibetu string 
   seibetu = "boy"
   serina := seibetu
   fmt.Println(serina)  // boy

   seibetu = "girl"
   fmt.Println(serina)  // boy

   fmt.Println("------------")

   var seibetu2 string 
   seibetu2 = "boy"
   serina2 := &seibetu2
   fmt.Println(*serina2)  // boy

   seibetu2 = "girl"
   fmt.Println(*serina2)  // girl
}
