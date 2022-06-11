## 今日のすること３つ
1. fizzbuzzの出力をtest.txtに書き出す
2. fizzbuzzのテストを構造体とforを使って書き出す
3. fizzbuzzの速度をPythonと比較して測る

## 1. fizzbuzzの出力をtest.txtに書き出す。
- 次のような構成にする
- (lib)はフォルダ
- (注意)libフォルダはpackage lib で作る
```go
fizzbuzz---main.go
         |-(lib)-lib.go
         |-go.mod
```
- main.goは以下のようにする
```go
package main

import (
	"fmt"
	"fizzbuzz/lib"
)

func main(){
	for i := 1; i < 1000; i++{
		n := lib.FizzBuzz(i)
		fmt.Println(n)
	}
}
```
## 2. fizzbuzzのテストを構造体とforを使って書き出す
```go
fizzbuzz---main.go
         |-(lib)---lib.go
         |       |-lib_test.go
         |
         |-go.mod
```
## 3. fizzbuzzの速度をPythonと比較して測る
- Goは`go build`で実行ファイル(fizzbuzz)を作る。
- Goは`time ./fizzbuzz`　Pythonは`time python3 fizzbuzz.py`
```python
for i in range(1000):
    if i % 15 == 0:
        print("fizzbuzz")
    elif i % 3 == 0:
        print("fizz")
    elif i % 5 == 0:
        print("buzz")
    else:
        print(i)
```
