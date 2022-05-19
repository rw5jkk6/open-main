## 変数の宣言の仕方２通り
```go
package main
 
import (
   "fmt"
)

func main(){
	var n int = 3  // 冗長
	m := 5         // 型推論
	fmt.Println(n * m)
}
```
- 1度の宣言で複数の変数を記述できる
```go
package main
 
import (
   "fmt"
)

func main(){
	var (
		x = 2
		y = 3
		z = 4
	)
	fmt.Println(x * y * z)
}
```
