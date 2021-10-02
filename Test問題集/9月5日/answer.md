Answer.0
```go:title
package main
 
import (
   "fmt"
)

type Person struct{
   name string
   age int
}

func (p *Person)add(){
   p.age++
}

func main(){
   p := Person{name: "suzuki", age: 20}
   p.add()
   fmt.Println(p.age)
}
```

Answer.1
```go:title
package main
 
import (
   "fmt"
)

type Vertex struct{
   X int
   Y int
}

func (v *Vertex)Scale(n int){
   v.X = v.X * n
   v.Y = v.Y * n
}

func main(){
   v := Vertex{X: 2, Y: 3}
   v.Scale(10)
   fmt.Printf("%+v\n", v)
}
```
Answer.2

```go:title
func Println2(a ...interface{})(n int, err error){
   return fmt.Fprintln(os.Stdout, a... )
}

Answer.3

package main

import (
   "fmt"
)

func main(){
   var m map[string]int
   m = make(map[string]int)
   m["nakaya"] = 3
   m["serinyan"] = 20

   // m := map[string]int{"nakaya":3, "serinyan":20}
   fmt.Println(m)
   m["nakaya"] = 4
   fmt.Println(m)
}
```
Answer.4



Answer.5
```go:title
func main(){
   var a interface{}
   a = "Hello"
   a = 1
   fmt.Println(a)
   b := a.(int)
   c := b + 2
   fmt.Println(c)

}
```
Answer.6
```go:title
func main(){
   name := []string{"satou", "suzuki", "yamada"}
   for _, v := range name{
      fmt.Println(v)
   }
}
```
Answer.7

Goプログラミング　cf160 list18-6

Answer.8
```go:title
type np int16
```
Answer.9
```go:title
func main(){
   for {
      fmt.Println("Hello serinyan")
   }
}
````
