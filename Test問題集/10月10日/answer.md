1. [ 数値は14です ]を2, 10, 16進数で出力する
2. 2の16乗をコマンドで2つでやってみる
3. 定数を使って、beast-modeは1.5として、[ヒトシのビーストモードは4.5です]と出力する
4. 複合リテラルとコンストラクタ関数の違いは
5. コンストラクタ関数を作る。以下を参考にする
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
12. ~複合リテラルとコンストラクタ関数の違いは?~
13. unix時間とは
14. unsignedとは
15. osパッケージを使ってカレントディレクトリを表示
16. mathパッケージを使って3と2の大きい方を表示
17. stringsパッケージを使って serina を大文字にする
18. strconvパッケージを使ってint 10を文字列にする
19. 文字列hogeをintにすることで、panicを起こす
20. var p *int　の出力は
21. User構造体を作って、name, plフィールドを作って "hitoshi","3"を入れる.そしてjsonに変換する
