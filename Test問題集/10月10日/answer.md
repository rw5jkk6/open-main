1. [ 数値は14です ]を2, 10, 16進数で出力する
```go:title
  fmt.Printf("%b\n %d\n %x\n", 14, 14, 14)
```
2. 2の16乗をコマンドで2つでやってみる 
```title:go
  bcコマンドを使う　　2^16
  コマンド置換え  echo $((2**16))
  追加  python3 2**16
```
4. 定数を使って、beast-modeは1.5として、[ヒトシのビーストモードは4.5です]と出力する
5. 複合リテラルとコンストラクタ関数の違いは
6. コンストラクタ関数を作る。以下を参考にする
```title:go
  type np struct{
      name string
      npl int
  }
```
6. 46 を４左シフト演算する
7. row文字列とは
8. 1110010を２進から16進へ
9. 16進のffffは何byte
10. コマンドライン引数に serina, akari, haruka, inu をランダムに出力して、[ ヒトシは ~ で妥協する]　とする
```title:go
package main

import (
	"fmt"
	"os"
   "math/rand"
   "time"
)

func main(){
   array := os.Args
   array2 := array[1:]
   rand.Seed(time.Now().UnixNano())
   fmt.Printf("hitoshiは%sで妥協する\n", array[rand.Intn(len(array2)+1)])
}
```
11. ~複合リテラルとコンストラクタ関数の違いは?~
12. unix時間とは
14. unsignedとは
15. osパッケージを使ってカレントディレクトリを表示
```title:go
   dir, _ := os.Getwd()
   fmt.Println(dir)
```
18. mathパッケージを使って3と2の大きい方を表示
```title:go
    fmt.Println(math.Max(3, 2))
```
20. stringsパッケージを使って serina を大文字にする
```title:go
    fmt.Println(strings.ToUpper("serina"))
```

22. strconvパッケージを使ってint 10を文字列にする
```title:go
        var num int64 = 10
	s := strconv.FormatInt(num, 10)
	fmt.Println(s)

	// defaultの数字はintなので、なぜ変換するのか不思議
	s1 := strconv.FormatInt(20, 10)
	fmt.Println(s1)
```

24. string 30　を int 30に変換してfmt.Printf()で出力する
```title:go
	v, _ := strconv.ParseInt("30",10,10)
	fmt.Println(v)
```
26. var p *int　の出力は     nil
27. User構造体を作って、name, plフィールドを作って "hitoshi","3"を入れる.そしてjsonに変換する
