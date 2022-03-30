●記号の問題（記述,７問）
複数呼び方がある場合は片方でもいい
⚪︎ ~
⚪︎ #
⚪︎ ^
⚪︎ { }
⚪︎ [ ]
⚪︎ `
⚪︎ |

●vimの問題（○×、３問）
⚪︎ノーマルモードにするのは、ESCと？     ctrl + [
⚪︎Helloを検索する                           :/Hello
⚪︎Helloをbyeに置き換える             :%s/Hello/bye/g

●前回の問題(○×、10問)
⚪︎ターミナルでカーソルを一番後ろに動かす    command + b
⚪︎ls /nakaya でエラーになるので error.txt に書き込む    ls /nakaya 2> error.txt
⚪︎file_a.txtをfile_c.txtに名前を変える     mv file_a file_c
⚪︎BeerをWhiskyに全部書き換える     sed "s/Beer/Whisky/g" file
⚪︎隠しファイルを含めて出力する       ls -a
⚪︎ロングフォーマットでファイルを出力する   ls -l
⚪︎拡張子が .txt だけを抜き出す      ls *.txt
⚪︎二列目を基準に降順で並べる         sort -r 2
⚪︎システムやミドルウェアに関する設定ファイルが格納されている      /etc
⚪︎ログファイルなどが格納されている    /var

●前の範囲(全問○×、15問)
1,
/dev/nullファイルを使う
⚪︎エラーメッセージだけを非出力する      ls / /nakaya 2> /dev/null

2,
dateコマンドを使う
⚪︎時:分で表示する          date "+%H:%M"

3,
catコマンドを使う。ここで作ったファイルは以降も使う
⚪︎Helloと書いたfileを作る。command + dで抜ける

cat > file.txt
Hello
[ctrl] + d
⚪︎Helloと書いたfileを作る。EOFで抜ける

cat << EOS >tile.txt
hello
EOS

4,
copy, moveコマンドを使う。(dir作成、移動後はcdコマンドは使わない)
⚪︎以下のディレクトリ、ファイルをつくる  mkdir dir; cd dir; mkdir dir2 dir3; touch fileA fileB
dir-|-dir2
      |-fileA
      |-fileB
      |-dir3
⚪︎dir2にfileA, fileBをコピーする  cp fileA fileB dir2
⚪︎dir3の下にdir2をコピーする   cp -r dir2 dir3
⚪︎dir/dir2にあるfileAをdirにfileCに名前を変えてコピーする   cp dir2/fileA ./fileC
⚪︎dirにあるfileAをfileDに名前を変える mv fileA fileD
⚪︎fileDをdir2に移動させる   mv fileD dir2
⚪︎dir3/dir2のfileAをdirに移動する  mv dir3/dir2/fileA .
⚪︎dirのディレクトリを全て表示して以下とあってるか確かめる　ls -R

5,
{}パラメータ展開を使う。
⚪︎パラメータ展開使わずにでtest6からtest9を作る          
for i in $(seq 6 9); do touch test$i.txt; done 
⚪︎パラメータ展開で以下のを内容を作って、http.txtファイルに保存する
http.hello.com
http.world.com
https.hello.com
https.world.com

touch  {http, https}.{hello, world}.com  | tr " " "\n"   > http.txt
⚪︎上で作ったtest5.txt~test8.txtのファイルを削除する     rm *.txt

●今回の範囲、コマンド編(15問)

1,（記述）
uniqコマンドを使う。以下のuniq.txtファイルを使う

japan
america
france
france
japan
england
america

⚪︎ユニークだけを表示する        
⚪︎重複しない行だけ表示する    
⚪︎fileの重複行を数えて、降順にする 

2,（記述）
lnコマンドを使う　testファイルを作る
⚪︎testファイルにtest_slinkというシンボリックリンクを貼る。そして消す  
⚪︎この時にtestファイルのリンクはいくつになる
⚪︎testファイルにtest_hlinkというハードリンクを貼る
⚪︎この時にtestファイルのリンクはいくつになる
⚪︎ハードリンクのlsコマンドによるファイルタイプは何になる
⚪︎シンボリックリンクにできてハードリンクにできないことは　　

3,(○×)
findコマンドを使う。
dir-|-dir2-|-fileA
      |          |-fileB
      |          |-fileD
      |
      |-dir3-dir4
      |-fileA
      |-fileB
      |-fileC
⚪︎fileAの場所を見つける    find . -name "fileA"    
⚪︎ディレクトリの数はいくつ       find . -type f | wc -l
⚪︎dir2の空のファイルを詳細表示する     find . -empty  -ls
⚪︎dir2の空のファイルを削除する          find . -empty -exec rm {} \;
<応用>
上とは関係ないfindコマンドを使う
⚪︎2日以上前、１年以内で上の空ファイルを見つけてくる 
find . -empty  -a -mtime +2 -and -mtime -365

4,（記述）
ターミナル内で計算する     1 + 1 = 2
⚪︎コマンド置き換えを使って  1 + 1 = 2 を出力する  
