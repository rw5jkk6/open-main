## codeを書く順番
1. 設計を考える
2. フォルダを作る `go mod init モジュール名`
3. コードを書く、テストも一緒に
4. コードを整形する  
5. `go mod tidy`
6. テストを実行する
7. 実行ファイルを作る

## ソースコードの整形(どっち使ってもいい)
- `go fmt`
- `gofmt -w ファイル名`
  - -w 修正する  
## library
- [標準ライブラリ](https://cs.opensource.google/go/go/+/refs/tags/go1.17.6:src/)
