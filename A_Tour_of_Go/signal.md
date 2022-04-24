## 動いているプロセスの止め方２つ
- CTRL + C でプロセスを止めることができる
- psコマンドで動いているプロセスを確認して、killコマンドでプロセス番号で止める


```go
package main

import (
	"fmt"
	"time"
)

func main(){
	for i := 0; i < 100; i++{
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}	
}
```
