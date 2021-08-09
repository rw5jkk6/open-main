package main
 
import (
   "fmt"
)

func main(){

   contory := map[int]string{1:"US", 81:"Japan", 86:"China"}
   c := contory[1]
   fmt.Println(c)        // US
   
   c, ok := contory[81]
   fmt.Println(c, ok)    // Japan  true

   contory[2]="Canada"
   fmt.Println(contory)   // map[1:US 2:Canada 81:Japan 86:China]

   delete(contory, 2)
   fmt.Println(contory)   // map[1:US 81:Japan 86:China]

   con2 := contory
   fmt.Println(con2)       // map[1:US 81:Japan 86:China]
   con2[3]="Mexico"
   fmt.Println(con2)       // map[1:US 3:Mexico 81:Japan 86:China]
   fmt.Println(contory)    // map[1:US 3:Mexico 81:Japan 86:China]
}
