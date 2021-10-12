- [1,2,3,4]
- 上の配列から以下を出力する
```title:go
1 2 3
1 2
1
```
- answer
```title:go
package main

import (
	"fmt"
)

func main(){
   a := []int{1,2,3,4}
   for i:=0; i < len(a); i++{
      fmt.Println("")
      for j:=0; j < len(a)-i-1; j++{
         fmt.Printf("%d ",a[j])
      }
   }
   fmt.Println("")
}
```
