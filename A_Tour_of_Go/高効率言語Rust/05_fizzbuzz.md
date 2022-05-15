```go
package main

import (
	"fmt"
)

func main(){
   for i:=1; i <= 1_000_000; i++{
      if i % 3 == 0{
         fmt.Println("fizz")
      }else if i % 5 == 0{
         fmt.Println("buzz")
      }else{
         fmt.Println(i)
      }
   }
}
```
