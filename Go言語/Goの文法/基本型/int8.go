package main
 
import (
   "fmt"
)

func main(){
   var a int8 = 3
   var aa int8 = -3
   fmt.Printf("%08b\n", a) // 00000011
   fmt.Printf("%08b\n", aa) // -0000011
   


   var i int8 = 127
   var ii int8 = -128
   var iii int8 = -127
   fmt.Printf("%08b\n", i) // 01111111
   fmt.Printf("%08b\n", ii) // -10000000
   fmt.Printf("%d\n", ii)  // -128
   fmt.Printf("%08b\n", iii)  // -1111111

}



