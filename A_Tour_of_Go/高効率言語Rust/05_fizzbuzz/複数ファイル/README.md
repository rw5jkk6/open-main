## fizzbuzz
- 関数名の頭は大文字にする
- 実行は `go run *.go`

## fizzbuzz_test
- テストファイル名は`ファイル名_test.go`とする
- テスト関数名は`Test_関数名`にする
- `go mod init ファイル名`
- `go mod tidy`
### テストの実行
- 全体の実行は `go test -v`
- 特定のテスト関数の実行 `go test -v -run Test_関数名`
