package main

import (
	"fmt"

	"golang.org/x/text/date"
)

type student_data struct{
   id int
   name string
   age int
}

func main(){
   id := []int{1, 2}
   name := []string{"satou", "tanaka"}
   age := []int{18, 20}
   
   for i:= 0; i < 2; i++{
      data := NewsetData(id[i], name[i], age[i])
   }
   fmt.Println(data)

}

// コンストラクタ関数
func NewsetData(id int, name string, age int) *student_data{
   d := new(student_data)
   d.id = id
   d.name = name
   d.age = age
   return d
}
