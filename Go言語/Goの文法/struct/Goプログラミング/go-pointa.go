package main

import (
	"fmt"
)

type person struct{
   name, superpower string
   age int
}

func main(){
   timmy := &person{
      name: "Timothy",
      age:10,
   }

   fmt.Println(timmy) // &{Timothy  10}
   
   // 本来はデリファレンスのため下の描き方だが、*を省略することができる
   // (*timmy).superpower = "test"

   timmy.superpower = "flying"
   fmt.Printf("%+v\n", timmy) // &{name:Timothy superpower:flying age:10}
}