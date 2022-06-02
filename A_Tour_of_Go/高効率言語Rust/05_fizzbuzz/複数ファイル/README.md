## fizzbuzz
- 関数名の頭は大文字にする
- 実行は `go run *.go`
- (注意) `go mod init ~`がなくても実行できるが赤線が出る

## fizzbuzz_test
- テストファイル名は`ファイル名_test.go`とする
- テスト関数名は`Test_関数名`にする。その関数の引数にはtesting Tをつける
- `go mod init ファイル名`
- `go mod tidy`
### テストの実行
- 全体の実行は `go test -v`
- 特定のテスト関数の実行 `go test -v -run Test_関数名`

## fizzbuzz_python
- 実行は `python3 -m unittest discover -v`
