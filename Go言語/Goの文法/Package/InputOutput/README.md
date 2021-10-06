- fileの入出力(上の方が原始的、下の方が高機能)
  - io.Read() <- 読み込むだけ。byte配列のbufferも自分で用意しないといけない
  - bufio.NewScanner()  <- 1行づつ読み込む
  - io/ioutil.ReadAll()  <- 全部読む
  - io/ioutil.ReadFile()  <- fileも開けて、全部読む

- ioで読み込まれたデータは全てbyteスライスになっているので、string()で文字列になる 
