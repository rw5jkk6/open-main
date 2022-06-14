## テストの作り方
- テストファイル名は`ファイル名_test.go`とする
- テスト関数名は`Test_関数名`にする。その関数の引数にはtesting Tをつける
- `go mod init ファイル名`
- `go mod tidy`
## テストの実行
- 全体の実行は `go test -v`
- 特定のテスト関数の実行 `go test -v -run Test_関数名`
- テストのカバレッジ
  - `-cover`をつける 
## mainファイルの実行
- `go run`は使うことができないので、`go build`で実行ファイルを作るとテストは実行されない
## Add
- main.goがないので本来はあり得ない形になっている

## fizzbuzz_test
- テストもmainパッケージになっているので本来はあり得ない形になっている
## swap_test
- これが正しい  
## fizzbuzz_python
- 実行は `python3 -m unittest discover -v`
