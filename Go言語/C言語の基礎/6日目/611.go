package main

import (
	"fmt"
)

type student_data struct{
   id int
   name string
   age int
}

func main(){
   data := []student_data{
      {1, "yamada", 18},
      {2, "satou", 19},
      {3, "oota", 18},
      {4, "nakata", 18},
   }

   for i := 0; i < 4; i++{
      fmt.Printf("学生番号:%d 名前:%s 年齢:%d\n", 
                  data[i].id, data[i].name, data[i].age)  
   }
}

// 学生番号:1 名前:yamada 年齢:18
// 学生番号:2 名前:satou 年齢:19
// 学生番号:3 名前:oota 年齢:18
// 学生番号:4 名前:nakata 年齢:18
