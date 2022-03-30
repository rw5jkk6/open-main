## 覚える単語(awkに関するもの)
- パターン(BEGINパターン、ENDパターンなど)
- アクション
- NF,NR
- / / <-正規表現
- ~ <-正規表現と使う
- フィールド(デフォルトはスペース)

## awk
> 1.3.d  
 `seq 5 | ` これ以降を書く
- 2,4 を正規表現で抜きだす
- 2,4 をc言語のような条件式で偶数を抽出
- アクションとパターンで書く
  - 2 偶数
  - 4 偶数
- アクションとパターンで書く
  - 1 奇数
  - 2 偶数
  - 3 奇数
  - 4 偶数
  - 5 奇数
- アクションとパターンで書く
  - 1 奇数
  - 2 偶数
  - 3 奇数
  - 4 偶数
  - 5 奇数
  - 合計 15
## 練習問題
  ```
  hitoshi,3,5
  serina,20,25
  sara,15,18
  kaoru,18,20
  cyurin,12,21
  ```
  - printで名前だけを出力する　`awk -F, '{print $1}' test.txt`
  - printfで名前だけを出力する `awk -F, '{printf("%s\n", $1)}' test.txt`
  - 最後の行だけを出力する `awk -F, '{print $NF}' test.txt`
  - 行数を表示する　　`awk -F, '{print NR}' test.txt`
  - 名前の頭にsがある人を抜き出す `awk -F, '/^s/' test.txt`
  - 名前の頭にsがある人を名前だけ抜き出す `awk -F, '/^s/{print $1}' test.txt`
  - 最初の行に name, normal, beast をつける `awk -F, 'BEGIN{print "name" "," "normal" "," "beast"}{print $0}' test.txt`
  - 最後の列の全部の行を平均する `awk -F, '{sum += $NF} END{print sum/NR}' test.txt`
  - normalを10以下の行を抜き出す `awk -F, '$3<10' test.txt`
  - normalを10以上16以下を抜き出す `awk -F, '$2 > 10 && $2 < 16' test.txt`
  - normalとbeastを合計したのを各行の最後の列に追加して出力する `awk -F, '{print $0, $2+$3}' test.txt`
