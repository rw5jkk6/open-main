// ./test hitoshi serina
// を実行して、hitoshi serina　を出力する

import (
	"fmt"
	"os"
)

func main(){
   // スライスの一部を切り取れる 
   arg := os.Args[1:]
   for _, v := range arg{
      fmt.Println(v)
   }
}
