- 九九の掛け算を出力する
```
1,  2,  3,  4,  5,  6,  7,  8,  9,
2,  4,  6,  8, 10, 12, 14, 16, 18,
3,  6,  9, 12, 15, 18, 21, 24, 27,
4,  8, 12, 16, 20, 24, 28, 32, 36,
5, 10, 15, 20, 25, 30, 35, 40, 45,
6, 12, 18, 24, 30, 36, 42, 48, 54,
7, 14, 21, 28, 35, 42, 49, 56, 63,
8, 16, 24, 32, 40, 48, 56, 64, 72,
9, 18, 27, 36, 45, 54, 63, 72, 81,
```

```go
package main
 
import (
   "fmt"
)

func main(){
    for i:=1; i < 10; i++{
	for j:=1; j < 10; j++{
	    fmt.Printf(" %2d,", i * j)
	}
	fmt.Println("")
    }
}
```
- 上の問題と似てるが、各行の最後のカンマをつけない(難易度高い)
```
1,  2,  3,  4,  5,  6,  7,  8,  9
2,  4,  6,  8, 10, 12, 14, 16, 18
3,  6,  9, 12, 15, 18, 21, 24, 27
4,  8, 12, 16, 20, 24, 28, 32, 36
5, 10, 15, 20, 25, 30, 35, 40, 45
6, 12, 18, 24, 30, 36, 42, 48, 54
7, 14, 21, 28, 35, 42, 49, 56, 63
8, 16, 24, 32, 40, 48, 56, 64, 72
9, 18, 27, 36, 45, 54, 63, 72, 81
```
```go
package main

import (
	"fmt"
	"strings"
)

func main(){

   for i:=1; i <= 9; i++{
      slist := []string{}
      for j:=1; j<=9; j++{
         s := fmt.Sprintf(" %2d",i*j)
         slist=append(slist, s)
      }
      s1 := strings.Join(slist, ",")
      fmt.Println(s1)
   }
}
```
