package main

import (
	"fmt"
)

type Person struct{
	name string
	age int
} 

func New(name string, age int) *Person{
	p := new(Person)
	p.name = name
	p.age = age
	return p
}

func main(){
	p := New("tarou", 10)
	fmt.Println(p)  // &{tarou 10}
	// フィールドは自動でデリファレンスされる
	fmt.Println(p.name) // tarou

}