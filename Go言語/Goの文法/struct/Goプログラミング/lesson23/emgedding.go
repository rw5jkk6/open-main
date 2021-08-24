package main

import (
	"fmt"
)

type length struct{
   normal int
   beast int
}

type size struct{
   length // length lengthでもOK
   thick int
}

type Pn struct{
   size
   color string
}

// func (p Pn)average() int{
//    return (p.size.length.normal + p.size.length.beast) / 2
// }

// size.lengthのところが省略されている。本では、これを転送と呼んでいる
func (p Pn)average() int{
   return (p.normal + p.beast) / 2
}

func main(){
   p := Pn{color: "white",
            size:size{thick: 3,
                      length: length{normal: 3, 
				     beast: 5},
            },
         }
   
   s := p.average()
   fmt.Println(s)
}
