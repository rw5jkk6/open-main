## 式と文
- 文(値を返さない。文はセミコロンで区切られる。改行すると自動でセミコロンがつく)
  - `i:=0` `num:=10` 
  - ```
    func add(a int,b int){
        fmt.Println(a + b)
    }
    ``` 
- 式(値を返す)
  - 変数や関数名 `num`  `add`
  - 定数 `1`  `0.1`
  - 文字列リテラル `"hello" `  
  - `i < 5` `i++`

- 具体例
```go
func main(){
	i:=5;
	fmt.Println(i);
};
```
## 文(statement)と式(expression)の違い
- 文は値を返さない (例) x=5
  - 文はセミコロンで終わる(Goでは省略してあることがある) 
- 式は値を返す   
  - 変数、関数名  (例) x, add()
  - 数字        (例) 3        
  - 文字列リテラル      (例) "hello"
  - 式を（）で囲んだもの   (例) (1+2)

## if文
- `if 条件式 {}`
- 条件式は`true`か`false`でしか判別できない

## switch式
- switch 簡易文; 式{  
  case 式:  
  }
- caseに式を書ける
  - 式のみ`case 1, 2`など
  - 変数を伴う式。この場合はswitchの横に式はいらない。if文と同じ 


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
