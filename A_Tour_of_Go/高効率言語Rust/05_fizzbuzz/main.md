- 1~100までの数字で３で割れるのをfizz、５で割れるのをbuzz、15で割れるのをfizzbuzz、それ以外は数字を表示する
```go
package main

import (
	"fmt"
)

func main(){
   for i:=1; i <= 100; i++{
      if i % 15 == 0{
         fmt.Println("fizzbuzz")
      }else if i % 3 == 0{
	 fmt.Println("fizz")
      }else if i % 5 == 0{
         fmt.Println("buzz")
      }else{
         fmt.Println(i)
      }
   }
}
```
- 1~50までの数を画面に表示し、3の倍数と３がつく整数の場合は数字ではなく『A』を表示する。なお、if文で else if と else を使わない。
```go
package main
 
import (
   "fmt"
)

func main(){
    for i:=1; i < 51;i++{
        if i % 3 == 0 || i % 10 == 3{
	    fmt.Println("A")
	　　　　　　　　continue
	}
	if i >= 30 && i <= 39{
	    fmt.Println("A")
	    continue
	} 
        fmt.Println(i)
    }
}
```
