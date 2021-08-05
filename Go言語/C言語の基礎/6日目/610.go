package main
 
import (
   "fmt"
)

type Student struct{
	id int
	name string
	age int
}

func main(){
	var data Student
	data.id = 1
	data.name = "山田太郎"
	data.age = 18
	fmt.Printf("学生番号:%d  名前:%s  年齢:%d\n", data.id, data.name, data.age)
}

// 学生番号:1  名前:山田太郎  年齢:18
