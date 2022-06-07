## 2つのモジュールの違い
- fizzbuzzはパッケージがmainの１つしかないが、culuc_multi_dirはパッケージが３つある。そのためculuc_multi_dirはmain.goはimportで他のパッケージから呼び出している。つまりパッケージ名が異なるとimportを使って、呼び出す必要がある

## fizzbuzz
- 関数名の頭は大文字にする
- 実行は `go run *.go`または`go run main.go`
- (注意) `go mod init ~`がなくても実行できるが赤線が出る

