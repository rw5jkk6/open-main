## 動いているプロセスの止め方２つ

- シグナルの種類を調べる方法 `man signal`
- シグナルとはプロセスに対して送られる信号
### *プロセス永久稼働しているのに対してシグナルを送る
***
1. CTRL + C でプロセスを止めることができる
2. psコマンドで動いているプロセスを確認して、killコマンドでプロセス番号で止める
```sh
ps
kill -KILL プロセス番号
```
***

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
### *標準入力
***
1. CTRL　+ d　でプロセスを止めることができる(これはEOFを入力している)
***
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
