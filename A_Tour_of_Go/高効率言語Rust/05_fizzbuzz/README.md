## 式と文
- 式
  - 変数や関数名 `num`  `add`
  - 定数 `1`  `0.1`
  - 文字列リテラル `"hello" `  
  - `i < 5` `i++`
- 文
  - `i:=0`

## for文を1行でなく、複数行で書く場合
```go
for i:=0
    i < 5
    i++{
    fmt.Println(i)
    }
```

## continueがあると、それ以降のプログラムを実行せずにfor文の初めに戻る
```go
package main

import (
	"fmt"
)

func main(){
   for i:=1; i <= 5; i++{
	   if i == 3{
		   fmt.Println(i)
		   //continue
	   } 
	   fmt.Println(i)
   }
}
```
## &&と||
- && は and
- || は or

```go
package main

import (
	"fmt"
)

func main(){
   for i:=1; i <= 5; i++{
   	   // ||を&&に変える 
	   if i % 2== 0 || i %4 == 0{
		   fmt.Println(i)
	   } 
   }
}
```

## 問題
1. for文で４から１２を出力する
2. for文で-10から10を2づつ偶数を出力する
3. for文で10から0まで1づつ減らすのを出力する  
