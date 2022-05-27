- 1~100までの数字で３で割れるのをfizz、５で割れるのをbuzz、15で割れるのをfizzbuzz、それ以外は数字を表示する。３つの方法で書いて、答えが合っているかdiffコマンドを使って確認する

- if文で書く

```go
package main

import (
	"fmt"
)

func main(){
   for i:=1; i <= 100; i++{
      if i % 15 == 0{
         fmt.Println("fizzbuzz")
      }else if i % 3 == 0{
	 fmt.Println("fizz")
      }else if i % 5 == 0{
         fmt.Println("buzz")
      }else{
         fmt.Println(i)
      }
   }
}
```
- switch文で書く
```go
package main
 
import (
   "fmt"
)

func main(){
	for i:=1; i <= 100; i++{
		switch {
	    //case i % 15 == 0:
	    case i % 3 == 0 && i % 5 == 0:
			fmt.Println("fizzbuzz")
	    case i % 3 == 0:
			fmt.Println("fizz")
	    case i % 5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}
```
- 無限forで書く
```go
package main
 
import (
   "fmt"
)

func main(){
	i:=1
	for{
		if i % 15 == 0{
			fmt.Println("fizzbuzz")
		}else if i % 3 == 0{
			fmt.Println("fizz")
		}else if i % 5 == 0{
			fmt.Println("buzz")
		}else{
			fmt.Println(i)
		}
		i++
		if i == 101{
			break
		}
	}
}
```
- constとifとswitchで書く
```go
package main

import (
	"fmt"
)

const (
	fizzbuzz int = iota
	fizz
	buzz
	num
)

func main() {
	for i := 1; i <= 100; i++ {
		var fb int
		if i%15 == 0 {
			fb = fizzbuzz
		} else if i%3 == 0 {
			fb = fizz
		} else if i%5 == 0 {
			fb = buzz
		} else {
			fb = num
		}
		switch fb {
		case fizzbuzz:
			fmt.Println("fizzbuzz")
		case fizz:
			fmt.Println("fizz")
		case buzz:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}

}

```
## 1~50までの数を画面に表示し、3の倍数と３がつく整数の場合は数字ではなく『A』を表示する。なお、if文で else if と else を使わない。
```go
package main
 
import (
   "fmt"
)

func main(){
    for i:=1; i < 51;i++{
        if i % 3 == 0 || i % 10 == 3{
	    fmt.Println("A")
	　　　　　　　　continue
	}
	if i >= 30 && i <= 39{
	    fmt.Println("A")
	    continue
	} 
        fmt.Println(i)
    }
}
```
