- cf50より
- 車は正数、新幹線は浮動小数で求める
```go
package main
 
import (
   "fmt"
)

func main(){
	const (
		moon = 384_400
		car = 80
		btrain = 300.0
	)
	fmt.Printf("月まで車で%d日\n", moon / car / 24)
	fmt.Printf("月まで新幹線で%f日\n", moon / btrain / 24)
}
```
