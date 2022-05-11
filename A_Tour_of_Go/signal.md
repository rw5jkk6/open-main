## 動いているプロセスの止め方２つ
- CTRL + C でプロセスを止めることができる
- psコマンドで動いているプロセスを確認して、killコマンドでプロセス番号で止める

- プロセス永久稼働
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
- 標準入力
```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	for {
		buffer := make([]byte, 1)
		_ , err := os.Stdin.Read(buffer)
		if err == io.EOF{
			fmt.Println("end")
			break
		}
		fmt.Printf("%s", string(buffer))
	}
}
```
